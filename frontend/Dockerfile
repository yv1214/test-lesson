FROM nginx:1.12-alpine
COPY ./enter.html /usr/share/nginx/html
COPY ./default.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
