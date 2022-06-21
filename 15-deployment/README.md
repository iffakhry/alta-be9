# Deployment

## Connect to EC2
```
ssh -i </directory/namafile.pem> <username-server>@<public-ipv4> 
```

## Transfer file using SCP
add -r if you want to transfer folder recursively

```
scp -i </directory/namafile.pem> </directory/namafile-transfer> <username-server>@<public-ipv4>:/home/ubuntu
```

## Update and upgrade Instance
```
sudo apt-get update
sudo apt-get upgrade
```

## Install Docker
```
sudo apt install docker.io
```