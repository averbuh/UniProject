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
This project aims to establish a robust and automated CI/CD pipeline for our software development process. By implementing continuous integration and continuous delivery (CI/CD), we aim to achieve significant gains in:

- **Efficiency:** Automating tasks like code building, testing, and deployment will free up developer time and reduce manual effort.
- **Quality:** Continuous integration with automated testing will help identify and fix bugs early in the development cycle.
- **Delivery Speed:** Frequent and automated deployments will allow for faster feature delivery and quicker response to user needs.
- **Reliability:** Consistent and automated processes will minimize human error and ensure a more reliable deployment process.
This document outlines the specific tools and procedures that will be implemented within the CI/CD pipeline. We'll detail the stages involved, from code checkout and security scans to testing, deployment, and monitoring. Additionally, we'll discuss the expected benefits and potential challenges associated with this project.

By successfully implementing a CI/CD pipeline, we can significantly improve our development workflow, deliver higher-quality software faster, and ultimately enhance user satisfaction.

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

# 🛤CICD
Continuous Integration and Continuous Deployment (CI/CD) is a set of practices and tools used by software development teams to automate the process of integrating code changes, testing them, and deploying them to production environments. 

CI/CD helps improve the efficiency and reliability of software development by automating the build, test, and deployment processes. This allows developers to quickly and easily merge their code changes into a shared repository, run automated tests to ensure the code is functioning correctly, and deploy the changes to production environments with minimal manual intervention.

By implementing CI/CD practices, teams can reduce the risk of introducing bugs and errors into their code, improve collaboration between team members, and accelerate the delivery of new features and updates to end users. Overall, CI/CD helps streamline the software development process and increase the overall quality of the code being produced.

- **CircleCI:** This is a continuous integration and continuous delivery (CI/CD) platform. It can automate many of the steps mentioned above, including running secret scans, linting, unit tests, and builds. CircleCI can be configured to trigger these tasks whenever there are code changes, allowing for faster feedback and earlier detection of issues.
![CICD pipeline](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---WHliGseN5cf9X37Ms63DV---figure---mfZQAvRtDj14EFcLMGcTLA.png "CICD pipeline")

## Pull Request to Dev
On this step run tests for smallest units of services like classes, functions and check for secrets in repository.

![Feature plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---02OMzReBvZH0wUDXVV70x---figure---3vHb4JfAk9hj0iBDfKpfYQ.png "Feature plus")



1. **Secret Scan:** This involves searching the codebase for any sensitive information that might have been accidentally committed, such as API keys, passwords, or credit card numbers. Secret scans help prevent such information leaks.
2. **Lint:** Linting is a static code analysis technique that checks for stylistic errors and potential bugs without actually running the code. It helps enforce coding conventions and identify areas for improvement.
3. **Code Quality with SonarCloud:**
    - SonarCloud can provide comprehensive analysis of code quality, including metrics for readability, maintainability, complexity, potential bugs, and code coverage. It can also track these metrics over time to identify trends and areas for improvement.
4. **Build:** The build process involves compiling the source code into an executable program or a deployable artifact. This typically involves steps like compiling, linking, and packaging the code along with any dependencies.
5. **Unit Test:** Unit tests are small, focused tests that verify the functionality of individual units of code, such as functions, classes, or modules. They help ensure that each piece of code works as expected in isolation.
6. **Code Coverage:** The percentage of code that is executed by the unit tests. High code coverage indicates that most of the code has been tested, reducing the risk of bugs.
## Push to Dev
### 


![Dev plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---slj6SJI47B9_OnlnqFoBo---figure---dt2_G6UMvJ2JnsUsnP1Xpw.png "Dev plus")



**2. SAST (Static Application Security Testing):**

- **Action:** Analyze the codebase for security vulnerabilities without executing it.
- **Tools:** SonarQube
**3. App Build:**

- Compile the source code into an executable program or a deployable artifact: Docker image, application build
**4. Upload Artifacts:**

- Transfer the built application files (DockerImage) to a repository for deployment.
**5. Deploy to Dev Environment:**

- Move the uploaded artifact to a development (Dev) environment for testing.
**6. Integration Tests, URL/API Checks:**

- Run automated tests that verify how different parts of application work together. URL/API checks ensure endpoints are reachable and functional.
- **Tools:** Postman (API testing)
**7. Functional Tests:**

- Run automated tests that verify the application's core functionalities from a user's perspective.
- **Tools:** Testing frameworks like Selenium (web UI)
**8. Automated UI Tests:**

- Run automated tests that interact with the application's user interface (UI) to ensure it behaves as expected.
- **Tools:** Testing frameworks like Selenium (web UI)


Quality assurance(QA) environment

- Helm values for qa:
    - app_mode: qa
    - app_url: https://qa.app.example.com
    - db_user: qa-user
    - db_password: qa-password
# 🛸Deployment
## Kubernetes and Backend
Kubernetes, often abbreviated as K 8 s, is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

Kubernetes allows developers to easily deploy and manage containerized applications across a cluster of machines. It provides features such as automated scaling, load balancing, self-healing, and rolling updates, making it easier to manage and scale applications in a dynamic and distributed environment.

With Kubernetes, developers can define their application's desired state using declarative configuration files, which specify the desired number of containers, resource requirements, networking rules, and other settings. Kubernetes then takes care of scheduling and managing the containers to ensure that the application runs smoothly and efficiently.

- ****
**Infrastructure as Code (IaC):** Terraform allows you to define your EKS cluster infrastructure as code. This code can be version controlled, shared with your team, and reused across environments. This promotes consistency, repeatability, and easier collaboration.



![image.png](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___s4vn9oIrPB5G1lN5l5juG.png "image.png")



### Network
- CNI Plugin - for Pod-to-Pod connection on different nodes( AWS CNI)
- Nginx Ingress controller - for routing, load balancing and more
- Services for static DNS names and ip adresses
### 
![K8s_worker](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---mYUbtYZEMqKwOB1afm1LC---figure---tbMByHI0RF5ynmZRu77Jlg.png "K8s_worker")

## Staging Environment
![Stage plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---6hOCpUon13noSdTmp-HXM---figure---9DH40NjOIIG2Jwn9B5E1Kg.png "Stage plus")

**Staging environments** consist of software, hardware, and configuration similar to the production environments. It is through these similarities testers can mimic the real-world production environment.

Staging environments are replicas of the production environments. It imitates the production environment as closely as possible to ensure application quality. The purpose of setting up a staging environment is to validate the application approaching the production stage to ensure the app will perform well post-deployment.

**Load testing** simulate high user loads on your application to assess its performance under stress. This helps identify bottlenecks and ensure the application scales effectively. 

- **Apache JMeter:** A free and powerful load testing tool with a large user base and extensive features for simulating various user behaviors.
**Dynamic Application Security Testing**. It's a security testing technique used to identify vulnerabilities in a running application. Unlike SAST (Static Application Security Testing) that analyzes code without execution, DAST simulates real-world attacks on a functioning application to see how it reacts.

- **OWASP ZAP:** Free and open-source, ZAP is a powerful and versatile DAST tool with a large community and extensive features.


Beta Test:

- **Identify Bugs and Issues:** Uncover any bugs, crashes, or functionality problems that the development team might have missed.
- **Evaluate Usability:** Assess how users interact with the product, identify any confusing or frustrating elements, and ensure it's user-friendly.
- **Gather Feedback:** Get valuable insights from real users about the product's features, performance, and overall user experience.


- Helm values for staging:
    - app_mode: stage
    - app_url: https://staging.app.example.com
    - db_user: staging-user
    - db_password: staging-password


## Production Environment


![Realease plus](/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---figure---SGGLMeq_k-JjmifqVQf1m---figure---3yj3XuJeOGQnwNcDixsycg.png "Realease plus")

- **Execution:**
    - **Deployment:** Move the application or service from the development/staging environment to the production environment using deployment tools and automation.
    - **Monitoring:** Closely monitor the application's performance and user activity after deployment to identify any problems.
- **Post-Deployment:**
    - **Analyze and Address Issues:** Evaluate any issues encountered and address them promptly.
    - **Gather Feedback:** Collect and analyze user feedback on the new version and use it for future improvements.
**Tools used in production rollouts:**

- **Monitoring Tools:** Tools like Prometheus, Grafana, or Datadog track application performance and health after deployment.


- Helm values for production:
    - app_mode: production
    - app_url: https://prod.app.example.com
    - db_user: production-user
    - db_password: production-password




# 😝Conclusion
The implementation of a CI/CD pipeline represents a significant step forward in our development process. By automating key tasks and integrating them into a continuous flow, we have established a foundation for faster, more reliable, and higher-quality software delivery.


<!-- eraser-additional-content -->
## Diagrams
<!-- eraser-additional-files -->
<a href="/README-cloud-architecture-1.eraserdiagram" data-element-id="5bdsd_EFGbYs8GSTeo1BM"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----6e3534aad42b3ca7c7fdff8216866ddb.png" alt="" data-element-id="5bdsd_EFGbYs8GSTeo1BM" /></a>
<a href="/README-cloud-architecture-2.eraserdiagram" data-element-id="vFy9TdHncMc-fAIb1BCsw"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----4a2240d1c071351a913135d20c926fd2.png" alt="" data-element-id="vFy9TdHncMc-fAIb1BCsw" /></a>
<a href="/README-cloud-architecture-3.eraserdiagram" data-element-id="On1Yi80bcZJI8_9qJc-cq"><img src="/.eraser/gLS3GlE11kV4Ul0K11od___XWe2mTYyEQVK1seoxfgZbWU7S5g1___---diagram----75d9043fe6e8ad3cc6d43bc65233591e.png" alt="" data-element-id="On1Yi80bcZJI8_9qJc-cq" /></a>
<!-- end-eraser-additional-files -->
<!-- end-eraser-additional-content -->
<!--- Eraser file: https://app.eraser.io/workspace/gLS3GlE11kV4Ul0K11od --->