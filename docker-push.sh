docker build -t kube-echo-headers .
docker tag kube-echo-headers:latest quay.io/ottintl/kube-echo-server:latest
docker push quay.io/ottintl/kube-echo-server:latest
