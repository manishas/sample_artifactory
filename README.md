# sample_artifactory
This is a sample project to demostrate the use of JFrog Artifactory integration in Shippable CI.

![image](https://cloud.githubusercontent.com/assets/4211715/21346909/68e35290-c6cc-11e6-96c2-61fd07633481.png)

Our sample project will download a data file named `data.json` from `pages` artifactory repository.
CI process will run a node.js program to template a static html file with information from `data.json`.
The result file `result.html` will be pushed to artifactory repository (`pages`).
It will upload the `result.html` to `pages/index.html` and a backup copy to `pages/backups/index.$BUILD_NUMBER.html`

![image](https://cloud.githubusercontent.com/assets/4211715/21347806/37984232-c6d0-11e6-90e6-3b043b0ac649.png)


### Step 1:
Add `shippable.yml` to your source control repository. Here is the `shippable.yml` file used in this sample project
```yml
language: none

build:
  pre_ci_boot:
    image_name: drydock/u14
    image_tag: tip
    pull: false
    options: '--privileged=true --net=bridge -e FOO=true -e BOO=false'

  ci:
    - jfrog rt dl pages/data.json
    - sudo apt-get install nodejs
    - npm install
    - node index.js
    - jfrog rt u result.html pages/index.html
    - jfrog rt u result.html pages/backups/index.$BUILD_NUMBER.html

integrations:
  hub:
    - integrationName: mycompany-jfrog
      type: artifactory

  notifications:
    - integrationName: mycompany-pipeline-trigger
      type: webhook
      payload:
        - versionName=$BUILD_NUMBER
      on_success: always
```

### Step 2: Create JFrog Artifactory Integration
In the yml, we could notice `mycompany-jfrog`, so we need to add a account Integration from Shippable UI with that name.
```yml
integrations:
  hub:
    - integrationName: mycompany-jfrog
      type: artifactory
```
Please refer to [Shippable documentation](http://docs.shippable.com/integrations/artifactRegistries/jfrogArtifactory/) for tutorial on how to add JFrog artifactory account integration.

### Step 3: Create Webhook Integration
This step is only needed, if you want to setup Continuous deployment after the CI passes. The following yml section is responsible for triggering deployments.
```
  notifications:
    - integrationName: mycompany-pipeline-trigger
      type: webhook
      payload:
        - versionName=$BUILD_NUMBER
      on_success: always
```
Please refer to [Shippable documentation](http://docs.shippable.com/integrations/notifications/webhooks/) for tutorial on how to add Webhook account integration.
