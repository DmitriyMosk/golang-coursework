#!/bin/bash

# Функция для проверки установленного пакета
is_installed() {
  command -v "$1" >/dev/null 2>&1
}

# Проверка и установка Docker
if is_installed docker; then
  echo "Docker уже установлен"
else
  echo "Docker не установлен. Установка Docker..."
  sudo apt-get update
  sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  sudo apt-get update
  sudo apt-get install -y docker-ce
  sudo systemctl start docker
  sudo systemctl enable docker
  echo "Docker установлен"
fi

# Проверка и установка kubectl
if is_installed kubectl; then
  echo "kubectl уже установлен"
else
  echo "kubectl не установлен. Установка kubectl..."
  curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
  sudo mv kubectl /usr/local/bin/
  sudo chmod +x /usr/local/bin/kubectl
  echo "kubectl установлен"
fi

echo "successful"