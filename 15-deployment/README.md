# Deployment

## Connect to EC2
```bash
ssh -i </directory/namafile.pem> <username-server>@<public-ipv4> 

# example
ssh -i file.pem ubuntu@54.25.5.56
```

## Transfer file using SCP
add -r if you want to transfer folder recursively

```bash
scp -i </directory/namafile.pem> </directory/namafile-transfer> <username-server>@<public-ipv4>:/home/ubuntu

# example
scp -i ~/file.pem file1.txt ubuntu@54.25.5.56:/home/ubuntu
```

## Update and upgrade Instance
```
sudo apt-get update
sudo apt-get upgrade
```

## Install Docker on ubuntu server
```
sudo apt install docker.io
```


## build Image
```bash
docker build -t <image-name>:<tag> .

# example
docker build -t api-images:latest .
```

## Show Image list
```
docker images
docker images list
```

## delete image
```
docker rmi <image-id>
docker rmi <image-name>
```

## Create Container from image
```bash
docker run -d 
-p <host-port>:<container-port> 
-e <env-name>=<env-value> 
-e <env-name>=<env-value> 
-v <host-volume>:<container-port> 
--name <container-name> <image-name>

# example: 
docker run -p 80:80 --name be9Container alta-be9-api:latest
```

## Show Container 
```bash
# show a running container 
docker ps

# show all container, including a stopped container
docker ps -a 
```

## Stop / Start Container
```
docker stop <container-name>
docker start <container-name>
```

## Remove Container
```
docker rm <container-name>
docker rm <container-id>
```

## Docker Logs Container
```
docker logs <container-name>
```

## Login Docker Hub and push to docker hub
```
docker login -u <username>

docker build -t <username-dockerhub>/<image-name>:<tag> .

docker push <username-dockerhub>/<image-name>
```

## Pull Image from container registry
```
docker pull <image-name>
```

## if you can't run docker in ubuntu (permission denied)
```
sudo usermod -a -G docker ubuntu

or

sudo chmod 777 /var/run/docker.sock
```


