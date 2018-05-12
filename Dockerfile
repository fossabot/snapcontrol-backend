FROM scratch

COPY bin/snapcontrol-backend /snapcontrol/snapcontrol-backend

CMD ["/snapcontrol/snapcontrol-backend"]
