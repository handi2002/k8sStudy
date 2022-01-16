FROM ubuntu
ENV MY_SERVICE_PORT=80
LABEL multi.author="Andy Han" multi.project="K8S study" other="geekbang"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
