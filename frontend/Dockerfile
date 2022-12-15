FROM node:12.16.1

WORKDIR /es-document-search-web/
COPY . .

RUN npm install
RUN npm run build

FROM nginx:1.15.3-alpine
LABEL MAINTAINER="wzj.7.29@gmail.com"


# 移除nginx容器的default.conf文件、nginx配置文件
RUN rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf
# 把主机的nginx.conf文件复制到nginx容器的/etc/nginx文件夹下
COPY .docker-compose/nginx.conf /etc/nginx/
COPY --from=0 /es-document-search-web/dist /usr/share/nginx/html

# 暴露2333端口
EXPOSE 2333