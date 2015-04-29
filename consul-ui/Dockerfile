FROM benschw/consul



ADD https://dl.bintray.com/mitchellh/consul/0.5.0_web_ui.zip /tmp/webui.zip
RUN mkdir /ui && cd /ui && unzip /tmp/webui.zip && rm /tmp/webui.zip && mv dist/* . && rm -rf dist



ENTRYPOINT ["/bin/consul"]
