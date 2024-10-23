pipeline {
    agent any

    stages {
        stage('SonarQube Analysis') {
            environment {
                SCANNER_HOME = tool 'SonarScanner'
                PROJECT_KEY = "devseconnect"
                SONAR_TOKEN = "sqa_3879f4ff886e948e1f3a433d6c554e84fdbb2164"
            }
            steps {
                withSonarQubeEnv('SonarQube') { 
                    sh '''
                    $SCANNER_HOME/bin/sonar-scanner \
                    -Dsonar.projectKey=$PROJECT_KEY \
                    -Dsonar.sources=. \
                    -Dsonar.host.url=http://sonarqube:9000
                    -Dsonar.login=$SONAR_TOKEN
                    '''
                }
            }
        }
        stage('Quality Gate') {
            steps {
                timeout(time: 1, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }
    }
}

