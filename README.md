<p><a target="_blank" href="https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>

# CICD pipeline for microservice application deployed in Kubernetes cluster in AWS


| Header | Questions | Solutions |
| ----- | ----- | ----- |
| Intoduction(Problem) | <ul><li>What's the problem/issue? </li><li>How does it bother us?</li></ul> | My solution(describe project in general) |
| Repository/CodeBase | <ul><li>What repostory might look like?</li></ul> | <ul><li>Frontend, Backend and Ops code</li><li>Branch strategy</li><li>Releases and tags</li></ul> |
| App Deployment(AWS,K8s) |  |  |
| CICD(Automatization process) |  |  |
| Сonclusion |  |  |
# ❓Introduction
## ?Tags:
- Github tags for releases in main branch 
- Helm charts and releases 
- Docker images and containers 
# 😺Codebase
Github repositories:

- For application team
    - Branches
    - One or many rHow does she bother us?epos for microservices
    - Environments
    - Tags
    - Microservices
- For operation team
    - Inside app repo or separate repo
    - Connection with app repo


## Branch strategy(Git Flow)
![image.png](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___V55rMXaM122TwX8rJpUv2.png "image.png")



Once a **develop** branch gets created out of the **master**, it will have the latest code of the addition function. Now, developers must work on their task, they will create a **feature** branch out of the develop branch. Feature branches are used by developers to develop new features for upcoming releases. It branches off from the develop branch and must merge into the develop branch back once the development of the new functionality completes.

**Hotfix** branches are required to take immediate action on the undesired status of the master branch. It must branch off from the master branch and, after fixing the bug, must merge into the master branch as well as the develop branch so that the current develop branch does not have that bug and can deploy smoothly in the next release cycle. Once the fixed code gets merged into the master branch, it gets a new minor version tag and is deployed to the production environment.

Now, once developer completes his development (subtraction and multiplication), he will then merge his code from the feature branch into the develop branch. But before he merges, he needs to raise a PR/MR. As the name implies, this requests the maintainer of the project to merge the new feature into the develop branch, after reviewing the code. All companies have their own requirements and policies enforced before a merge. Some basic requirements to get a feature merged into the develop branch are that the feature branch should get built successfully without any failures and must have passed a code quality scan.

Once dev's code is accepted and merged into the develop branch, then the release process will take place. In the **release** process, the develop branch code gets merged to the release branch. The release branch basically supports preparation for a new production release. The code in the release branch gets deployed to an environment that is similar to a production environment. That environment is known as staging (pre-production). A staging environment not only tests the application functionality but also tests the load on the server in case traffic increases. If any bugs are found during the test, then these bugs need to be fixed in the release branch itself and merged back into the develop branch.

Once all the bugs are fixed and testing is successful, the release branch code will get merged into the master branch, then tagged and deployed to the production environment.

# 🛸App Deployment(AWS,K8s)
## Kubernetes and Backend
Kubernetes, often abbreviated as K 8 s, is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

Kubernetes allows developers to easily deploy and manage containerized applications across a cluster of machines. It provides features such as automated scaling, load balancing, self-healing, and rolling updates, making it easier to manage and scale applications in a dynamic and distributed environment.

With Kubernetes, developers can define their application's desired state using declarative configuration files, which specify the desired number of containers, resource requirements, networking rules, and other settings. Kubernetes then takes care of scheduling and managing the containers to ensure that the application runs smoothly and efficiently.



[﻿K8s_worker](https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od?elements=tbMByHI0RF5ynmZRu77Jlg)K8s Worker Node:

![K8s_worker](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---WpqanK55TwkDAPEk68e7Z---figure---tbMByHI0RF5ynmZRu77Jlg.png "K8s_worker")

### Network
- CNI Plugin - for Pod-to-Pod connection on different nodes( AWS CNI)
- Nginx Ingress controller - for routing, load balancing and more
- Services for static DNS names and ip adresses
- **DNS server** add-on [﻿https://coredns.io/](https://coredns.io/)  #what
### Configs






### Helm releases




### Session state data
![[12 factor app summary#^bm87qv]]

Micro-services should be stateless. Store all session data to datastores like Redis or Memcached.

### **Port binding** - Export services via port binding
In a local development environment, the developer visits a service URL like `http://localhost:5000/` to access the service exported by their app. In deployment, a routing layer handles routing requests from a public-facing hostname to the port-bound web processes.

### Logs
- [﻿Fluent](https://www.fluentd.org/) 
Grafana - Loki - Fluentd

### Todo
- [ ] Create **Helm** chart for app:
    - [x] DB
    - [ ] Monitoring
        - [ ] Add and setting Prometheus
        - [x] Graphana
    - [ ] Services
    - [ ] App deployment
    - [ ] Secrets (passwords and more)
    - [ ] Configuration files
## AWS (Terraform & Terrastruct)
Terraform is an open-source infrastructure as code tool created by HashiCorp. It allows users to define and provision infrastructure resources using a declarative configuration language. With Terraform, you can define your infrastructure in code using configuration files, called Terraform scripts, which describe the desired state of your infrastructure. Terraform then automates the process of provisioning and managing infrastructure resources across various cloud providers, such as AWS, Azure, Google Cloud, and more.

Terraform provides a way to manage infrastructure resources in a consistent and repeatable manner, making it easier to scale and maintain infrastructure configurations. It also supports infrastructure drift detection, allowing you to compare the actual state of your infrastructure with the desired state defined in your Terraform scripts.

- [ ] Write **Terraform** module:
    - [ ] EKS


# 🛤CICD
Continuous Integration and Continuous Deployment (CI/CD) is a set of practices and tools used by software development teams to automate the process of integrating code changes, testing them, and deploying them to production environments. 

CI/CD helps improve the efficiency and reliability of software development by automating the build, test, and deployment processes. This allows developers to quickly and easily merge their code changes into a shared repository, run automated tests to ensure the code is functioning correctly, and deploy the changes to production environments with minimal manual intervention.

By implementing CI/CD practices, teams can reduce the risk of introducing bugs and errors into their code, improve collaboration between team members, and accelerate the delivery of new features and updates to end users. Overall, CI/CD helps streamline the software development process and increase the overall quality of the code being produced.



![CICD pipeline](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---HdsBwYhKN2jONxpciao-P---figure---mfZQAvRtDj14EFcLMGcTLA.png "CICD pipeline")

## Pull Request to Dev
![Feature plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---z0423zcaEkxMDBBSNqMbl---figure---3vHb4JfAk9hj0iBDfKpfYQ.png "Feature plus")

## Push to Dev
### Tests:


![Dev plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---5TjPSH0RvnheWKQdvXhWY---figure---dt2_G6UMvJ2JnsUsnP1Xpw.png "Dev plus")

## Staging Environment
![Stage plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---F6gdue-x8u-sKlNFdyppN---figure---9DH40NjOIIG2Jwn9B5E1Kg.png "Stage plus")

## Production Environment


![Realease plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---jyXDMTGgkvc28u6fw6ktT---figure---3yj3XuJeOGQnwNcDixsycg.png "Realease plus")





# 😝Conclusion



<!-- eraser-additional-content -->
## Diagrams
<!-- eraser-additional-files -->
<a href="/README-CI/CD Pipeline with CircleCI and AWS Deployment-1.eraserdiagram" data-element-id="YUt59gsIKI9glFJueyQ4y"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----76ec49a498c9a2c947a50df130266cc0-CI-CD-Pipeline-with-CircleCI-and-AWS-Deployment.png" alt="" data-element-id="YUt59gsIKI9glFJueyQ4y" /></a>
<a href="/README-cloud-architecture-2.eraserdiagram" data-element-id="5bdsd_EFGbYs8GSTeo1BM"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----6e3534aad42b3ca7c7fdff8216866ddb.png" alt="" data-element-id="5bdsd_EFGbYs8GSTeo1BM" /></a>
<a href="/README-cloud-architecture-3.eraserdiagram" data-element-id="vFy9TdHncMc-fAIb1BCsw"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----4a2240d1c071351a913135d20c926fd2.png" alt="" data-element-id="vFy9TdHncMc-fAIb1BCsw" /></a>
<a href="/README-cloud-architecture-4.eraserdiagram" data-element-id="On1Yi80bcZJI8_9qJc-cq"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----c647751ffe8bf737b1ac9e88bd23ad7f.png" alt="" data-element-id="On1Yi80bcZJI8_9qJc-cq" /></a>
<!-- end-eraser-additional-files -->
<!-- end-eraser-additional-content -->
<!--- Eraser file: https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od --->