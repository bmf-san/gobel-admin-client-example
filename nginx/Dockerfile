FROM nginx:1.19.0-alpine

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./conf.d/gobel-admin-client-example.conf /etc/nginx/conf.d/goble-admin-client-example.conf

RUN ln -sf /dev/stdout /var/log/nginx/access_gobel_admin_client_example.log \
    && ln -sf /dev/stderr /var/log/nginx/error_gobel_admin_client_example.log \