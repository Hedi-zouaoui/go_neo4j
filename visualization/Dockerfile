
FROM nginx:alpine

COPY index.html /usr/share/nginx/html
COPY test.json /usr/share/nginx/html

EXPOSE 7070


CMD ["nginx", "-g", "daemon off;"]