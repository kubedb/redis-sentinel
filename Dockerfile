FROM redis:6.2.1
COPY peer-finder /usr/local/bin/

RUN chmod +x /usr/local/bin/peer-finder

COPY on-start.sh /on-start.sh

RUN chmod +x /on-start.sh

