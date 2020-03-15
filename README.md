#greetrpc serve a testare le potenzialità di gRPC

[![Known Vulnerabilities](https://snyk.io/test/github/axamon/greetgrpc/badge.svg?targetFile=Gopkg.lock)](https://snyk.io/test/github/axamon/greetgrpc?targetFile=Gopkg.lock)

Il client e il server possono dialogare su ip e porta cusomizzabile.
Il protoccolo di trasmissione è http2 con compressione gzip e per cifrare i dati si usano certificati TLS.

Questo è un esempio di utilizzo di gRPC unary ma è possibile implementare anche il full duplex.

May API con serializzazione json "REST" in peace, gRPC sta arrivando! 