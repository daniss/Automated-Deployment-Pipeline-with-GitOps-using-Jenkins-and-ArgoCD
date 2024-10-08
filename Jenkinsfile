pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/daniss/Automated-Deployment-Pipeline-with-GitOps-using-Jenkins-and-ArgoCD'
                echo 'Checking out code..'
            }
        }
        stage('Build') {
            steps {
                cd 'go-api'
                sh 'go mod tidy'
                sh 'go build -o go-api'
                cd '..'
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Docker Build') {
            steps {
                def buildname = "danios149/go-api"
                sh "docker build -t ${buildname}:latest ."
                echo 'Building Docker Image..'
            }
        }
        stage('Docker Push') {
            steps {
                def buildname = "danios149/go-api"
                docker.withRegistry('https://index.docker.io/v1/', 'dockerhub') {
                    sh "docker push ${buildname}:latest"
                }
                echo 'Deploying....'
            }
        }
        stage('Deploy to minikube') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'

                sh 'kubectl rollout status deployment/go-api'
                echo 'Deploying to Kubernetes..'
            }
        }
    }
}