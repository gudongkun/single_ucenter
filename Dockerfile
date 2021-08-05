FROM alpine
ADD enlight_ucenter /enlight_ucenter
ENTRYPOINT [ "/enlight_ucenter" ]
