# Use an official Nginx image as the base image
FROM nginx:alpine

# Copy your HTML file to the Nginx HTML directory
COPY index.html /usr/share/nginx/html
COPY test.json /usr/share/nginx/html
# Expose the default HTTP port
EXPOSE 7070

# Start Nginx to serve the HTML file
CMD ["nginx", "-g", "daemon off;"]