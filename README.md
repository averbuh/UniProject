<p><a target="_blank" href="https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>

## Streamlining Deployments: Microservice CI/CD Pipeline for EKS
# Introduction
This project aims to establish a robust and automated CI/CD pipeline for software development process. By implementing continuous integration and continuous delivery (CI/CD). We aim to achieve significant gains in:

- **Efficiency:** Automating tasks like code building, testing, and deployment will free up developer time and reduce manual effort.
- **Quality:** Continuous integration with automated testing will help identify and fix bugs early in the development cycle.
- **Delivery Speed:** Frequent and automated deployments will allow for faster feature delivery and quicker response to user needs.
- **Reliability:** Consistent and automated processes will minimize human error and ensure a more reliable deployment process.
This document outlines the specific tools and procedures that will be implemented within the CI/CD pipeline. We'll detail the stages involved, from code checkout and security scans to testing, deployment, and monitoring. Additionally, we'll discuss the expected benefits and potential challenges associated with this project.

By successfully implementing a CI/CD pipeline, we can significantly improve our development workflow, deliver higher-quality software faster, and ultimately enhance user satisfaction.

### Project's tools and services:
**Services:**

- **Frontend**: Web UI with Vue framework
- **Backend**: 3 microservices using REST API for communication  
- **Database**: Postges
**Build & Test Tools in this project:**

- **Go tools (built-in):** `go build` , `go test`  for building and testing your Go applications.
- **Linters:** Tools like go fmt, go vet, or staticcheck can help maintain code quality and catch potential issues early.
**CI/CD Pipeline:**

- **CircleCI:** A CI/CD platform that will run your build, test, and deployment pipelines.
**Infrastructure Provisioning:**

- **Terraform:** To define and provision your infrastructure on AWS, including EKS clusters and supporting resources.
**Deployment & Configuration Management:**

- **Helm:** A package manager for Kubernetes that will manage the deployment and configuration of Go microservices as Helm charts. We can also install public charts like Graphana, Prometheus, Postges DB, Redis. 
**Container Orchestration:**

- **Kubernetes (kubernetes.io):** Runs your containerized Go microservices within the provisioned EKS cluster.
**Cloud Provider:**

- **AWS:** Provides the cloud infrastructure and services, including Amazon EKS for managing your Kubernetes cluster.
- **Amazon EKS:** A managed Kubernetes service on AWS that simplifies cluster management.
**Additional Considerations:**

- **Container Registry:** Container registry Amazon ECR (Elastic Container Registry) to store Docker images for microservices.
- **Monitoring & Logging:** Prometheus and Grafana for monitoring microservices and Fluentbit for logging.
## Branch strategy(Git Flow)
A branch strategy is a set of guidelines that define how developers use branches in a version control system (VCS) like Git. It helps manage code changes, avoid conflicts, and maintain a clear version history, especially when multiple developers are working on the same project.

![image.png](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___V55rMXaM122TwX8rJpUv2.png "image.png")



Once a **develop** branch gets created out of the **master**, it will have the latest code of the addition function. Now, developers must work on their task, they will create a **feature** branch out of the develop branch. Feature branches are used by developers to develop new features for upcoming releases. It branches off from the develop branch and must merge into the develop branch back once the development of the new functionality completes.

**Hotfix** branches are required to take immediate action on the undesired status of the master branch. It must branch off from the master branch and, after fixing the bug, must merge into the master branch as well as the develop branch so that the current develop branch does not have that bug and can deploy smoothly in the next release cycle. Once the fixed code gets merged into the master branch, it gets a new minor version tag and is deployed to the production environment.

Now, once developer completes his development (subtraction and multiplication), he will then merge his code from the feature branch into the develop branch. But before he merges, he needs to raise a PR/MR. As the name implies, this requests the maintainer of the project to merge the new feature into the develop branch, after reviewing the code. All companies have their own requirements and policies enforced before a merge. Some basic requirements to get a feature merged into the develop branch are that the feature branch should get built successfully without any failures and must have passed a code quality scan.

Once dev's code is accepted and merged into the develop branch, then the release process will take place. In the **release** process, the develop branch code gets merged to the release branch. The release branch basically supports preparation for a new production release. The code in the release branch gets deployed to an environment that is similar to a production environment. That environment is known as staging (pre-production). A staging environment not only tests the application functionality but also tests the load on the server in case traffic increases. If any bugs are found during the test, then these bugs need to be fixed in the release branch itself and merged back into the develop branch.

Once all the bugs are fixed and testing is successful, the release branch code will get merged into the master branch, then tagged and deployed to the production environment.

# 🛤CICD
This section outlines the Continuous Integration and Continuous Delivery (CI/CD) pipeline designed to streamline the development and deployment process for our project. Our CI/CD pipeline leverages a suite of tools to automate key stages, including:

- **Building and Testing:** Code changes will trigger automated builds and comprehensive tests to ensure code quality and functionality before deployment.
- **Infrastructure Provisioning:** Infrastructure on AWS, including the EKS cluster, will be automatically provisioned and configured using Terraform.
- **Deployment and Configuration Management:** Containerized microservices will be deployed and managed on the EKS cluster using Helm charts.
The big picture of CICD pipeline from Pull Request to Production Deploy:

![CICD pipeline](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---Aq7IIOuFMiHbdAnl1HTBd---figure---mfZQAvRtDj14EFcLMGcTLA.png "CICD pipeline")



Now, we will look on pipeline in details. It have 4 separate parts:

- Developer create Pull Request to Dev branch
- Merge Feature branch in Dev ( Deploy to dev environment )
- Deploy to Stage environment
- Deploy to Production environment
Pipeline results will trigger real-time updates in our dedicated Slack channel, keeping everyone informed of success or any potential issues. These messages helps developers immediately understand the root of the problem.

## Pull Request to Dev
This stage meticulously tests the building blocks of our code (classes and functions) through **unit testing**. Additionally, a thorough security scan is conducted to identify any sensitive information (secrets) accidentally left within the repository.

![Feature plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---9Xh2xAYz1c3ksT_U-uLLc---figure---3vHb4JfAk9hj0iBDfKpfYQ.png "Feature plus")



1. **Secret Scan:** This involves searching the codebase for any sensitive information that might have been accidentally committed, such as API keys, passwords, or credit card numbers. Secret scans help prevent such information leaks.
2. **Lint:** Linting is a static code analysis technique that checks for stylistic errors and potential bugs without actually running the code. It helps enforce coding conventions and identify areas for improvement.
3. **Code Quality with SonarCloud:**
    - SonarCloud can provide comprehensive analysis of code quality, including metrics for readability, maintainability, complexity, potential bugs, and code coverage. It can also track these metrics over time to identify trends and areas for improvement.
4. **Build:** The build process involves compiling the source code into an executable program or a deployable artifact. This typically involves steps like compiling, linking, and packaging the code along with any dependencies.
5. **Unit Test:** Unit tests are small, focused tests that verify the functionality of individual units of code, such as functions, classes, or modules. They help ensure that each piece of code works as expected in isolation.
6. **Code Coverage:** The percentage of code that is executed by the unit tests. High code coverage indicates that most of the code has been tested, reducing the risk of bugs.
## Push to Dev
Once the code is deemed complete, the developer submits a pull request for the feature branch to be merged into the `dev` branch. This initiates a review process, potentially involving the QA team or automated testing pipelines in the QA environment, to ensure quality and functionality before the merge is approved.



QA stands for **Quality Assurance**. It's a broad term encompassing the processes and activities that ensure a product or service meets the specified requirements for quality.

- Developers play a crucial role in QA, writing unit tests and ensuring code quality.


![Dev plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---37HXVPmBT4VUZ0lAXZpV8---figure---dt2_G6UMvJ2JnsUsnP1Xpw.png "Dev plus")



**SAST (Static Application Security Testing):**

- Analyze the codebase for security vulnerabilities without executing it.
- **Tools:** SonarQube
**App Build:**

- Compile the source code into an executable program or a deployable artifact: Docker image, application build
**Upload Artifacts:**

- Transfer the built application files (DockerImage, Helm chart) to a repository for deployment.
** Deploy to Dev Environment:**

- Move the uploaded artifact to a development (Dev) environment for testing.
** Integration Tests, URL/API Checks:**

- Run automated tests that verify how different parts of application work together. URL/API checks ensure endpoints are reachable and functional.
- **Tools:** Postman (API testing)
** Functional Tests:**

- Run automated tests that verify the application's core functionalities from a user's perspective.
- **Tools:** Testing frameworks like Selenium (web UI)
** Automated UI Tests:**

- Run automated tests that interact with the application's user interface (UI) to ensure it behaves as expected.
- **Tools:** Testing frameworks like Selenium (web UI)


Quality assurance(QA) environment

- Helm values for qa:
    - app_mode: qa
    - app_url: https://qa.app.example.com
    - db_user: qa-user
    - db_password: qa-password
# Environments
## Kubernetes Cluster
Kubernetes, often abbreviated as K8s, is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

Kubernetes allows developers to easily deploy and manage containerized applications across a cluster of machines. It provides features such as automated scaling, load balancing, self-healing, and rolling updates, making it easier to manage and scale applications in a dynamic and distributed environment.

With Kubernetes, developers can define their application's desired state using declarative configuration files, which specify the desired number of containers, resource requirements, networking rules, and other settings. Kubernetes then takes care of scheduling and managing the containers to ensure that the application runs smoothly and efficiently.

**Infrastructure as Code (IaC):** Terraform allows you to define your EKS cluster infrastructure as code. This code can be version controlled, shared with your team, and reused across environments. This promotes consistency, repeatability, and easier collaboration.

**Components of an EKS Cluster:**

- **Control Plane:** Manages the cluster and consists of API server, scheduler, and controller manager. EKS manages the control plane for you.
- **Worker Nodes:** Virtual machines or containers that run your containerized applications. You can manage worker nodes manually or use auto scaling groups in EKS.
- **Pods:** The smallest deployable unit in Kubernetes, consisting of one or more containers and shared storage.
- **Deployments:** Manage the lifecycle of your applications by scaling and updating Pods.
![image.png](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___s4vn9oIrPB5G1lN5l5juG.png "image.png")

### Worker Node:
![K8s_worker](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---zDBlenvOWYBHPx2-Eh0kF---figure---tbMByHI0RF5ynmZRu77Jlg.png "K8s_worker")



## Helm chart
A Helm chart is a package that contains all the resources needed to deploy an application to Kubernetes, such as:

- Kubernetes deployment manifest (defines how to run containerized applications)
- Service manifest (defines how to expose the application)
- ConfigMaps and Secrets (store sensitive configuration data)
- Additional resource manifests (e.g., Ingress resources for external access)
When you install a Helm chart using the `helm install` command, Helm creates a release on your Kubernetes cluster. This release represents a running instance of the application defined in the chart. Each release has a unique name that you specify during installation.

### Helm charts for deploying three microservices
**Chart Structure:**

- **Chart.yaml:** This file defines the metadata for the chart, including name, version, dependencies on other charts (if any), and a brief description.
- **templates/**: This directory contains templates used to generate Kubernetes manifests (Deployment, Service, etc.) for your microservices. These templates leverage placeholders that will be filled with values during deployment.
- **values.yaml:** This file defines the default values for all the placeholders used in the templates. You can override these defaults when installing the chart to customize the deployment for your specific environment.
**Microservice Deployment:**

The templates directory will likely contain separate subdirectories for each microservice, for example:

- `templates/restourant/` : This directory might contain:
    - `deployment.yaml.tpl` : A template for the Deployment resource that defines how to run your first microservice container image. It will use placeholders for things like image name, container port, and resource requests/limits.
    - `service.yaml.tpl` : A template for the Service resource that exposes the first microservice through a load balancer or NodePort for external access. Placeholders will define the service port and type.
- Similar directories (`templates/recipts/`  and `templates/suppliers/` ) with templates for deployments and services for the other two microservices.
**Values Customization:**

The `values.yaml` file might include default values for:

- `image.repository` : Base name of the container image repository for all microservices 
- `image.tag` : Default image tag (e.g., latest).
    - **Git Commit SHA:** Use Git commit SHAs as tags to tie deployments to specific code commits in the version control system.
- `service.port` : Default port for all microservice services.
- Individual service ports and resource requirements can be defined for each microservice.
### Upgrade Helm release on EKS cluster
**Helm Upgrade:** 

Replacing `<release-name>` with the actual name of Helm release and `<chart-name>` with the name of the Helm chart:

```
helm upgrade <release-name> <chart-name> [flags] [values.yaml]
```
**Values Overrides:** If specific configuration changes are required, utilize the `--set`  flag or a values file within the `helm upgrade`  command to override default values in the Helm chart.

## Staging Environment
Invoked when the Dev environment pipeline passes all checks.

![Stage plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---Vdpul0s9bX3jDXeTUVLHA---figure---9DH40NjOIIG2Jwn9B5E1Kg.png "Stage plus")

**Staging environments** consist of software, hardware, and configuration similar to the production environments. It is through these similarities testers can mimic the real-world production environment.

Staging environments are replicas of the production environments. It imitates the production environment as closely as possible to ensure application quality. The purpose of setting up a staging environment is to validate the application approaching the production stage to ensure the app will perform well post-deployment.

**Load testing** simulate high user loads on your application to assess its performance under stress. This helps identify bottlenecks and ensure the application scales effectively. 

Tools like** Apache JMeter:** A free and powerful load testing tool with a large user base and extensive features for simulating various user behaviors.

**Dynamic Application Security Testing**. It's a security testing technique used to identify vulnerabilities in a running application. Unlike SAST (Static Application Security Testing) that analyzes code without execution, DAST simulates real-world attacks on a functioning application to see how it reacts.

Tools like** OWASP ZAP:** Free and open-source, ZAP is a powerful and versatile DAST tool with a large community and extensive features.

**Beta Test:**

- **Identify Bugs and Issues:** Uncover any bugs, crashes, or functionality problems that the development team might have missed.
- **Evaluate Usability:** Assess how users interact with the product, identify any confusing or frustrating elements, and ensure it's user-friendly.
- **Gather Feedback:** Get valuable insights from real users about the product's features, performance, and overall user experience.


- Helm values for staging:
    - app_mode: stage
    - app_url: https://staging.app.example.com
    - db_user: staging-user
    - db_password: staging-password


## Production Environment
Initiated after successful completion of the Staging pipeline.

![Realease plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---CSNU4tPKA4aOfMigB2nry---figure---3yj3XuJeOGQnwNcDixsycg.png "Realease plus")

- **Execution:**
    - **Deployment:** Serve application from the staging environment to the production environment using Helm.
    - **Monitoring:** Closely monitor the application's performance and user activity after deployment to identify any problems.
- **Post-Deployment:**
    - **Analyze and Address Issues:** Evaluate any issues encountered and address them promptly.
    - **Gather Feedback:** Collect and analyze user feedback on the new version and use it for future improvements.
        - **Monitoring Tools:** Tools like Prometheus, Grafana, Fluentbit track application performance and health after deployment.
**Rollback strategy**

A rollback strategy is a crucial plan of action that outlines the steps to take in case a deployment to a production environment causes unexpected issues or problems. Its primary goal is to quickly revert to a previously working state, minimizing downtime and impact on users.

- **Rollback:** Reverts to a previously working state, essentially undoing the deployment.  Helm allows rollbacks to previous releases using the `helm rollback` command. 
- **Rollforward:** Involves deploying a new version that specifically addresses the issue that caused the initial deployment failure. This approach can be considered if a simple rollback isn't feasible or desirable.


- Helm values for production:
    - app_mode: production
    - app_url: https://prod.app.example.com
    - db_user: production-user
    - db_password: production-password
# Conclusion
The implementation of a CI/CD pipeline represents a significant step forward in our development process. By automating key tasks and integrating them into a continuous flow, we have established a foundation for faster, more reliable, and higher-quality software delivery.


<!-- eraser-additional-content -->
## Diagrams
<!-- eraser-additional-files -->
<a href="/README-cloud-architecture-1.eraserdiagram" data-element-id="On1Yi80bcZJI8_9qJc-cq"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----7c326c8e6ff824b2825349bdd3eabd00.png" alt="" data-element-id="On1Yi80bcZJI8_9qJc-cq" /></a>
<!-- end-eraser-additional-files -->
<!-- end-eraser-additional-content -->
<!--- Eraser file: https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od --->