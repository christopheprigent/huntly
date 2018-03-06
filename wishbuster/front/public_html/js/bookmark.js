(function($) {
    var setCookie = function(name, value, days, path) {
            /*Valeur par défaut de l'expiration*/
            var expires = '';
            /*Si on a spécifié un nombre de jour on le convertit en dae*/
            if (days != undefined && days != 0) {
                var date = new Date();
                /*On évite les dates négatives*/
                if (days < 0) {
                    date.setTime(0);
                } else {
                    date.setTime(date.getTime() + Math.ceil(days * 86400 * 1000));
                }
                expires = '; expires=' + date.toGMTString();
            }
            /*Si on a pas spécifié de path on pose le cookie sur tout le domain*/
            path = path || '/';
            document.cookie = name + '=' + encodeURIComponent(value) + expires + '; path=' + path;
        },
        getCookie = function(c_name) {
            var i, x, y, ARRcookies = document.cookie.split(";");
            for (i = 0; i < ARRcookies.length; i++) {
                x = ARRcookies[i].substr(0, ARRcookies[i].indexOf("="));
                y = ARRcookies[i].substr(ARRcookies[i].indexOf("=") + 1);
                x = x.replace(/^\s+|\s+$/g, "");
                if (x == c_name) {
                    return unescape(y);
                }
            }
            return undefined;
        },
        getMemberID = function() {
            /* hack hors : https://secure.fr.vente-privee.com/ */
            if (getCookie("infoMember") === undefined)
                setCookie("infoMember", 18294862);

            return getCookie("infoMember");
        };
    // $(document).ready(function() {
    var apikey = getMemberID(),
        host = 'http://hunt.ly:8000',
        huntBookMarkHref = "javascript:void((function(){var hunt=document.createElement('script');";
    huntBookMarkHref += "hunt.setAttribute('src','" + host + "/components/hunt.js?api_key=";
    huntBookMarkHref += apikey + "'); document.querySelector('body').appendChild(hunt);})());";
    $("#huntlyBookMark").attr("href", huntBookMarkHref);
    // });
})(window.jQuery);