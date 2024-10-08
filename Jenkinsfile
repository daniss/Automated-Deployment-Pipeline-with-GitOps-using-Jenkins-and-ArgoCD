pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/daniss/Automated-Deployment-Pipeline-with-GitOps-using-Jenkins-and-ArgoCD'
                echo 'Checking out code...'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }
        stage('Docker Build') {
            steps {
                docker.withRegistry('https://index.docker.io/v2/', credentialsId: 'dockerhub') {
                    docker.build("danios149/go-api").push('latest')
                }
                
                echo 'Building Docker Image...'
            }
        }
        stage('Deploy to Minikube') {
            steps {
                script {
                    sh 'kubectl apply -f k8s/deployment.yaml'
                    sh 'kubectl apply -f k8s/service.yaml'
                    sh 'kubectl rollout status deployment/go-api'
                }
                echo 'Deploying to Kubernetes...'
            }
        }
    }
}
