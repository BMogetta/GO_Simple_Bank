# **A Platform-Agnostic SEcurity TOken [PASETO](https://paseto.io/)**
Paseto is a cryptographically secure, compact, and URL-safe representation of claims intended for space-constrained environments such as HTTP Cookies, HTTP Authorization headers, and URI query parameters.

Key Differences between Paseto and JWT
Unlike JSON Web Tokens (JWT), which gives developers more than enough rope with which to hang themselves, Paseto only allows secure operations. JWT gives you "algorithm agility", Paseto gives you "versioned protocols". It's incredibly unlikely that you'll be able to use Paseto in an insecure way.

Caution: Neither JWT nor Paseto were designed for stateless session management. Paseto is suitable for tamper-proof cookies, but cannot prevent replay attacks by itself.

`go get -u github.com/o1egl/paseto`