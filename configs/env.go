package configs

import "os"

func IsProductionEnvironment(IsProduction bool) {
	//////////////////////// Handling Production Environment ////////////////////////

	if IsProduction == true {
		os.Setenv("GIN_MODE", "release")
		os.Setenv("PORT", "8000")
		os.Setenv("SENDGRID_API_KEY", "")
		os.Setenv("ISSUER", "")
		os.Setenv("AUDIENCE", "entitlementregistry-dev.endpoints.neuralways.cloud.goog")
		os.Setenv("PRIVATE_KEY", "")                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               //this should be in an env file
		os.Setenv("ORGREG_ACCESS_PRIVATE", "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCt3G3YSLdWPqYv\nPtc8x+jo74Hmwd9w9lpOG8clIFJoaR4bnn4/1XABl+iUqB/9d+DnVf1G5GXz7Dvv\nT8hKYWRB9+eS+EQXAHlgRyzhS6ffo/KWnznEEhuWRwPBKUtZzVWvmitAFWug83rB\noS0ERcDoM/pbmG+82ocuIKbiFlMtCWPCngypuw74s56nREO7g3DNbeqEx8A9SUMP\nTPKu616loFnXcyV4yVzMPNfadY7hH4NSDazmpRV77+6eLCR7DC8OAm0xZUih5ynu\ns8050JUDKXC0jfiFQBGh601Jrza5l26zR9dHaoTDa7DKr7ecbtwhJBXQwX42PM4X\nI+bwkpErAgMBAAECggEACWSWoqVlvPPBiAfob6Y1P8XAyjYV25p+TSipslMetfld\nNmIdBqyB/UlhuAulPs0Sii0BHqq0P2C3PuSb444uG8pIPo50efGn8mWmAmioodJ6\nXAeE2Befr0IYQn/39tq7Ewj3O+JsYCkHVsgIZeoG1AS1ipafuB+Y8gsIORVlBP0d\nZLPEPJbuo9Ngwsjiy3tiwR63sWRVBwX141uzj4L9t1b/GTeAX6UueYz6ZSt3GHcJ\nVDfaF02NBSrZ5jlSm1TFHPUwhq3Xd19zaawlQ6qNSh9nZ9UV1jRShk2aSldORb1L\ncMv+MYW8n2oKL7vKG9NBi6UtsWpDrWhH5Wj6gnzgUQKBgQDcWyiqrPT7BIjEZvUf\nSRiuQhfXZcKUju7yh2VNl4ydcBLpGhrfESzltkc6jIYtDij4MIJq6jADUA3PYA2s\nQwbqSJJfzZX/s1KBhQJi+E2lQmC3gbbguN1cm3SP6k7xPkyjT6T3i1M/5O5+UNMP\nHTcnkPyKfKzsO8f+6Bc/sB/8TQKBgQDJ++7IlVgvbBegNFUaAjz/0gnPDWBi6zVf\nz9oIowTqAyFKvbw6IHMAvtmc9NHpX2nnx9ny8rK++KdudPsT3Q+A5gxogDjidqLq\nwPhlmAwMxR7MPs6VzRjbfkjPBLzSoy+3VKNhex6ufhXt8wsKLoaDPFPLDnrfOLAZ\niIuTGACfVwKBgAD35ZXabLktq+wxB6K7j4A0iiZILDK67hHixm8w201V0L91fei9\nCpV315hKRcys40Ez9V7C6LMbBvSCV40KPl2JNBt3q9ET2Kj5DKmirbSsNgCYcVVL\neq2NgVWgkfg9bdjJ4M+1rQf62ZF0iXjnh/aEjJePNLNVD+vu2StOetWBAoGBAMl6\nieUBvBaJz0NaEG3CP9C7Vld4M2ZY00QAlLA61OMNPj9M5h5hUfYzaeT7/Qu6xmyS\n2SBT4eNHK9St0eotf5KHZ2cySAv2ZWqgFOdadOKkGqWIHRAmpx4DaWM/iY7fgvBv\n+A17/831LbDO2HtOj5UHL1qE1iImbNr7bqbKsYw/AoGACS0pfZIFhLrZQ0xqIefS\nlJBQOBVbHPt4NNG/NCmyW5tgogfvO9JMJ86uqC1R97lQLbXNsMzO57A+9128Bg6J\nLWHZCCMXQsNfGjWQa9cClS2LsXewcvnsphr7WxkjjtvsM9Vc3OLFV5Qo2KkQ3Kwg\ne+GIJsfRp5TCn0MBQmLUQEc=\n-----END PRIVATE KEY-----\n") //this should be in an env file
		os.Setenv("ORGREG_Issuer", "witmeg-prod-orgregistry@neuralways.iam.gserviceaccount.com")
		os.Setenv("ORGREG_Audience", "orgregistry-prod.endpoints.neuralways.cloud.goog")
		os.Setenv("ORGREG_BASEURL", "https://orgregistry-prod.endpoints.neuralways.cloud.goog:5000")

	}
	//////////////////////// Handling Development Environment ////////////////////////

	if IsProduction == false {
		os.Setenv("ISSUER", "witmeg-entitlement-dev@neuralways.iam.gserviceaccount.com")
		os.Setenv("AUDIENCE", "entitlementregistry-dev.endpoints.neuralways.cloud.goog")
		os.Setenv("PRIVATE_KEY", "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCxzwo/Uw8O2+g0\nkYNZjDT2wWgaUMpvMR5DLvT2dqY96IxzFpSDt0pMoLWAfvvHt+GvdFKyjUTQ79yM\ng8NSpAl9SYEU7eq9qin/el/ZtlqWcSQXbjLLcPnSwwsvpdardyd6gIjG9PLDCsJm\njJ66zxW+E+jAJhVhfj434absRDsh92+25RpogxEI2gSjIXmc2tmhGo96UtWMP5Dl\nQffcvjqd4ZimXPy46Op45d9iLR9G5IdQw//uexihhL642VCHV5x626owF3OSB8uk\nWEGj/U/SKxnLMDFm9y/o7Ui88x3YI/+mgzB566Lp2FQX0PdJROqoXNGoLeu0LX5L\npGfeu5StAgMBAAECggEAF7GMHwYVYpiBJ8ggJaWeRSdelOshHpIAA1nVgKYs3uEH\nHirrHWYmO3HK3dWN/X4OT9ADWRGkEudzjzHbrD55mT1bOwu9FQsAz69h6GP5TETF\nV23yzNNAJwqgiffiVDTriNJp54EqfD0JzJp+aExCDHMgt8IbZUNqTQ3/rUomSw8x\ngpUBAuqFmEn499GlCgPja4yIJfDXItZV4TLTcGYFX1/dwhBqzoNOd1mTPstbJ/Cr\nNuQ/MpflnLTELwnX8nAmgfMjKafzHuK0o7XFT//Qi85bC1Pvw3d0cgBMrU9ulA79\n16v/mQlqRMQZZ+oxM5cNWbICeqPK/I46TGCK9jcpYQKBgQDxf3AXQVvy3cfj3Gh8\nh4lYtl5wFHReqoQOU9NhJ/tQQYKg0RzpDUWB/IhBF0pDOZG7/wM1iqsl1Zc6FTxD\nUg/LulEQwrMq8hO1kgRYizdAnAqPEW4GVon/uOh3NIW/uMKE5vGlk1pYRW8CG3lH\n6FJdDT3Wf4n4r7sXTAzoNOzBtQKBgQC8fIGKlh7OlgV08XHRcc508aeGEkP6EmFK\nH3AnKM1qqDMUvR6iHBV4S9IB1RWiD2Fq3Asg5oVliylweguYcRuVAc4Tu1WpmcX1\nsldl88aKsYrW1ZP+70QQKchLWCxLK9oUoLwZRzc6zqvqtjSmvMf2JbygjyTTM4xh\nACdGPnxCGQKBgQCQOT5U9RiXo3G0GDZiYINa8pbVsrGIs6bn85HijQEeffEoX5o3\naFChf5JT5TfZi/EH6BUfjbAGiVemPXZe8gwy7BNrwgtEU4wHvKtRwdp+gHIh3RPr\nPxVtHh3eywQxaoLo4g9ghe7Hlrfj90pn8kv1/Gl4x3Q7i59dG3L93lWinQKBgAgp\nhBqSQadEi9vykLE3Mm+R/2Uo6VSs8xWYouyvMaYKvpuJTQhMGwnojqGRrqwj6l4M\nRVcZZHFvQ04hmtrxGvaq61ejZyQ/CnmHTY2WruRWp+KBwTTq+5GPnaLVoKVosrrt\n8vdNRDEBviQd283UovwfyjlRuZCDzPJuZUT/c7YRAoGBAM9LqvNUUh5bgu17kx//\nePbzfeGWDoEb4y2yrQVSZ5sk0zQCLnaekXnRvkwXMcVZ5MQS7eH7Z6Ve+vkNzMH1\nBwA+xDzLlLFPS/ZLiIiDYoTH0OqUvDSvWF060Z9snoiw2O0bjSL3/rO11/vlDZGh\n3HP/0444paT3q2sam4jglj5E\n-----END PRIVATE KEY-----\n")           //this should be in an env file
		os.Setenv("ORGREG_ACCESS_PRIVATE", "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCh1ZeraxLFRIdd\nKRDg9keZ+KWKpPYoSH5JBZZqLCFpzvJ2Nc0n3qsfgh+EebMtIZ23RNcd/XFndjt4\nM0QBROva/bf6JINUNWAaDqLyzjdxE67nngXKO7KsJywrxys5dJsCD4X3CeRCfHMB\nChSxqO6znruYSP6wZ7BYOzSh/Tia+ImGZW/Y/4ox5RsTaYs12/OTFG5O/AQJyF42\nGrUQ8eiV5A39mt/QhIKy2kxR+3U1qYr+pYLBusZTJ5pmf3Z8bPApiCrJTrIM02n1\nbA5krWYq7LvH0xfPpJp2yNKTWP2q06GDSBdguh7vCTH26MKMuVsgEapqr9+Xgb5w\nQZIsrLplAgMBAAECggEAJiViXYe80i5byOAF6Dzs9jRXkuVc13R0/fHY+VhDHP1o\nhegbHk98Xw4tLyz/iL8QfbM2kI1dpoFGNchBupncDU+6RrZGFXqKV0IFcz2b8GWv\nF83V+AwXyr/mB8Cn5eBgLD1r+qXEfc53TjKp25GiCEk9V0E3dW3PcZCtEJTPO6nD\nFHxJnrTVTYWMmxD6WtimsF7C9XcQwsWLJNqMpCPwgBCm6jb0FrPdkVzFfUeB5z6p\nQOzXncOG+9SZzjEFELUwPd1rkyXuvjDCJyYIC8oOVz6aXFg63ia2Vr7tKMMXcK1L\ndSH81IeJu70opT5ZMKMtTh9STxh7vxvoGhj1XrmUYQKBgQDZ/C9OOaZTgtfWZ9J3\n9WECs64m0cEEGQHQ5erceTkI7pFF2GeMFM7iHEVCk/g68CLRf7qq8SeO7heC753e\nzaqfPx6y2LnathWWKqGo9JZSnfSao7sgAwpKnAwzkvk5gdzGkgOm/VTXl0p1oL+I\nC/ZbdIcQ39XE8AiqhhPG8hDdoQKBgQC+Dpc5kr8LUTpRLRLysNnGB5EUSld0S5GJ\nE58Qn6JKvvLH1Sjyr2Bkn3FMgoxXXRpeX4yTosSH0R5/rkjHJSyVtqJelFAE/Q3K\nLpNoOqO+3e/SLSLjuFeZ/nyJ7VMW51blYC9BqLsa1YYtabygBONl3VaQmmmQ+xe6\n/Dh7u/Q+RQKBgE7nE5BzGZtc3xeCmRJGxYie1UjD3+PF0Xx13J9VWam8AOWRQVnZ\nJlF9/ivtN4rN5SyD2WOKytDXqbcBWNmPBFWEGW6POCbMhd/HKqDnd4oMFr/qyA78\nSrta5Ke+aWvBb9KIg73sMMugvZj6SJnMCq5r7X/UJaKVZ+2hfoMil6xBAoGBALM8\ns2V6fW1A54ySw6HZIyfdDvdhvORSOChcNtDIm5d+SRwobPuiZXe/3mvXWl8x0KJb\n3gmFoW6syuSDCuO0rT1sBejIWka6F/0qu526X56rU3aVVqgHu6XhGLZC+O9KsOqV\nMTvs0rfV0zqU02zd706r1O1+kjxGcPTI1MPsV8U5AoGABfL6nxSmn3NzRH4kEnSX\nAkhMP/vHyhN51AnretiH17hDIW5TM1z8SpDMv8DOWny+qRFhqJibFzAP5qV1J8Vk\nS2m0fGVD0JAXw0/dGjbijk+iu+nFIKR2WRfqWUqKNwo9X8nCvCQeDkhogWMHvw3W\nTtmlRwNmRe7h+PaAxlKAPPM=\n-----END PRIVATE KEY-----\n") //this should be in an env file
		os.Setenv("ORGREG_Issuer", "witmeg-dev@neuralways.iam.gserviceaccount.com")
		os.Setenv("ORGREG_Audience", "orgregistry-dev.endpoints.neuralways.cloud.goog")
		os.Setenv("ORGREG_BASEURL", "https://orgregistry-dev.endpoints.neuralways.cloud.goog:5000")
		os.Setenv("ORGREG_PATHURL", "/v1/organisation/add")
	}
}