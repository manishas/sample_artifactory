# sample_artifactory (Pipelines)
This sample project demostrates continuous deployment of artifacts from JFrog Artifactory, with help of Shippable's declarative yml pipelines.


### Step 1: Create JFrog Artifactory Integration
This integration will be used to pull the artifact from JFrog and having a `file` resource in Shippable. In our sample project, we have named the integration as `mycompany-jfrog`.

Please refer to [Shippable documentation](http://docs.shippable.com/integrations/artifactRegistries/jfrogArtifactory/) for tutorial on how to add JFrog artifactory account integration.

### Step 2: Create Node Cluster Integration
This integration will point to the nodes of a cluster, where the artifacts could be deployed. In our sample project, we have named the integration as `mycompany-node-cluster`.

Please refer to [Shippable documentation](http://docs.shippable.com/integrations/deploy/nodeCluster/) for tutorial on how to add Node Cluster account integration.

### Step 3: Create resources
In Shippable pipelines, the file named `shippable.resources.yml` could be used to define the resources. 

We will be defining two resources here i.e. [file](http://docs.shippable.com/pipelines/resources/file/) and [cluster](http://docs.shippable.com/pipelines/resources/cluster/)

```yml
resources:
  - name: app-package
    type: file
    integration: mycompany-jfrog
    pointer:
      sourceName: pages/index.html
    seed:
      versionName: initial

  - name: app-cluster-2
    type: cluster
    integration: mycompany-node-cluster
```

### Step 4: Create Jobs
In Shippable pipelines, the file named `shippable.jobs.yml` could be used to define the jobs.

```yml
jobs:
  - name: app-manifest
    type: manifest
    steps:
      - IN: app-package
      - IN: app-deploy-trigger

  - name: app-deploy
    type: deploy
    steps:
      - IN: app-manifest
        force: true
      - IN: app-cluster-2
```

Please refer [manifest](http://docs.shippable.com/pipelines/jobs/manifest/) and [deploy](http://docs.shippable.com/pipelines/jobs/deploy/) to know how those jobs could be used in different ways

### Step 5: Create Triggers
Triggers are used to trigger the jobs. They can be defined in `shippable.triggers.yml`.

Please refer [here](http://docs.shippable.com/pipelines/triggers/) for detailed information about them.

### Step 6: Setup Pipelines
Now, the source repository with the above ymls could be used as `sync` repository to set up pipelines. More information about this step could be obtained from [here](http://docs.shippable.com/pipelines/gettingStarted/).

After the sync process is complete. You will be able to see the following SPOG in Pipelines tab.
![image](https://cloud.githubusercontent.com/assets/4211715/21351746/2c56d424-c6e4-11e6-85e7-241d4e4b503e.png)

Whenever the `app-deploy` job runs, it fetches the file artifact `pages/index.html` ( mentioned in `app-package` ) from JFrog Artifactory and transfers it to the nodes given in the Node cluster integration.

All the artifacts transferred from Shippable to the nodes will be inside `/tmp/shippable` folder.

