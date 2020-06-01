pipeline{
    agent any
    tools {
        go 'go-1.14.3'
        terraform 'Terraform'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages{
        stage("Clone"){
            steps{
                sh 'terraform --version'
                echo "========== Cloning the git repository: " + env.BRANCH_NAME
                git branch: 'master', url: 'https://github.com/vocacorg/terraform-provider-template.git'
                echo "Content in working directory"
                sh "ls -la ."
            }
            post{
                success{
                    echo "======== Repository cloned successfully ========"
                }
                failure{
                    echo "======== Unable to clone the repository ========"
                }
            }
        }
        stage("Code Quality"){
            environment {
                scannerHome = tool 'SonarqubeScanner'
            }
            steps{
                echo "========== Running Sonarqube quality checks"
                withSonarQubeEnv('sonarqube') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
                timeout(time: 10, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true, credentialsId: 'webhook-secret'
                }
            }
            post{
                success{
                    echo "======== Sonarqube quality checks completed ========"
                }
                failure{
                    echo "======== Failed sonarqube quality checks ========"
                }
            }
        }
        stage("Code Coverage"){
            steps{
                echo "========== Estimating the code coverage"
                sh "go test -coverprofile=coverage.out ./..."
                sh "mkdir coverage"
                sh "go tool cover -html=coverage.out -o coverage/index.html"
                publishHTML (target : [allowMissing: false,
                    alwaysLinkToLastBuild: true,
                    keepAll: true,
                    reportDir: 'coverage',
                    reportFiles: 'index.html',
                    reportName: 'Code Coverage'])
            }
            post{
                success{
                    echo "======== Code coverage report generated successfully ========"
                }
                failure{
                    echo "======== Unable to generate the code coverage report ========"
                }
            }
        }
        stage("Integration Testcases") {
            when {
                // Only run test cases when the branch name is 'staging'
                expression { env.BRANCH_NAME == 'staging' }
            }
            steps{
                echo "========== Running integration testcases"
                sh "go test -coverprofile=coverage.out ./..."
                sh "mkdir coverage"
                sh "go tool cover -html=coverage.out -o coverage/index.html"
                publishHTML (target : [allowMissing: false,
                    alwaysLinkToLastBuild: true,
                    keepAll: true,
                    reportDir: 'coverage',
                    reportFiles: 'index.html',
                    reportName: 'Code Coverage'])
            }
            post{
                success{
                    echo "======== Integration testcases completed ========"
                }
                failure{
                    echo "======== Unable to run integration testcases ========"
                }
            }
        }
        stage("Build"){
            steps{
                echo "========== Building the repository"
                sh 'go build'
                sh "ls -la ."
            }
            post{
                success{
                    echo "======== Build completed successfully ========"
                }
                failure{
                    echo "======== Unable to build the code ========"
                }
            }
        }
        stage("Run Terraform"){
            steps{
                echo "========= Running terraform files"
                sh 'terraform init'
            }
            post{
                success{
                    echo "======== Runnint terraform testcases ========"
                }
                failure{
                    echo "======== Unable to run terraform testcases ========"
                }
            }
        }
    }
    post{
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}