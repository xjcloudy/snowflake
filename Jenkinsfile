pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/xjcloudy/snowflake.git', branch: 'master')
      }
    }
    stage('build') {
      parallel {
        stage('build') {
          steps {
            sh 'echo "step1"'
            sh 'echo "step2"'
          }
        }
        stage('') {
          steps {
            sh 'echo "step3"'
          }
        }
      }
    }
  }
}