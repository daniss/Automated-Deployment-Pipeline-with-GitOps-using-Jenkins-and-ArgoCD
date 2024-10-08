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
                script {
                    cd 'go-api'
                    def buildname = "danios149/go-api"
                    docker.withRegistry('https://index.docker.io/v1/', 'fcb556ee-f565-41f1-9006-0cc7916f7711') {
                        docker.build(buildname).push('latest')
                    }
                    cd '..'
                }
                echo 'Building Docker Image...'
            }
        }
        stage('Docker Push') {
            steps {
                script {
                    cd 'go-api'
                    def buildname = "danios149/go-api"
                    docker.withRegistry('https://index.docker.io/v1/', 'fcb556ee-f565-41f1-9006-0cc7916f7711') {
                        sh "docker push ${buildname}:latest"
                    }
                    cd '..'
                }
                echo 'Pushing Docker Image...'
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
