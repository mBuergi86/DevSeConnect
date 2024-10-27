pipeline {
    agent any

    stages {
      stage('SCM Checkout') {
        steps{
          git branch: 'main', url: 'https://github.com/mBuergi86/DevSeConnect.git'
        }
      }
      stage('SonarQube Analysis') {
        steps {
          script {
            def scannerHome = tool name: 'sonar_scanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
            withSonarQubeEnv('SonarQube') {
              sh "${scannerHome}/bin/sonar-scanner" \
              "-Dsonar.projectKey=devseconnect" \
              "-Dsonar.sources=." \
              "-Dsonar.host.url=http://sonarqube:9000" \
              "-Dsonar.login=sqp_093d42af62a29c56f20da4ce238f3bd4096dae7a"
            }
          }
        }
      }
    }
}

