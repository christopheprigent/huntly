//https://neeed.com/bm/neeed.js.php?api_key=xiuhr897Cue
(function() {
    /** parametrage **/
    var borderSize = 4, // bien penser a  changer le css aussi
        iframeWidth = 200,
        imageWidth = 150,
        imageHeigth = 150,
        imageMaxWidth = 1855,
        imageMaxHeigth = 1855,
        url = 'https://neeed.com/bm/',
        _prices = ['[itemprop="price"]', '#priceblock_ourprice', '#price', '#product-price', '#actualPriceValue', '#our_price_display', '#prezzo-yard', '#product_view_price_new', '#productPrices', '#Prixprod2', '#blocQuantitePrix', '#discounted_price', '#priceblock_ourprice', '#product_price', '#pp_price', '#price-field', '#shop-price', '#prixFinal', '#apparel-price', '#hidden_price', '#with_old_price', '#selected_price', '#listing-price', '#buy_btn', '#product-page-body-right-column-price', '#prix_pas_promotion_euro_fiche_produit', '#prodpage-buy-price', '#productPrice', '#design-price', '#now_price', '#pricereg', '#s-price', '#first_block', '.prix.tcc', '.price .sale', '.price-sales', '.price .px_boutique', '.price ins', '.price', '.currentPrice', '.price_article', '.prix', '.regular-price', '.productSpecialPrice', '.currency_EUR', '.ourprice', '.price_output', '.prodprice', '.lePrix', '.prixProduit', '.price_block', '.darty_prix', '.userPrice', '.packagePrice', '.prices', '.gshpProductPrice', '.Price_Productinfo', '.product_view_price', '.goodPrice', '.good_spantotal', '.product-price', '.det_price', '.mp_current_price', '.product_price', '.prijs', '.p-price', '.price-box', '.price-sales', '.prd-amount', '.sglprice', '.entry_price', '.prix_fiche', '.ProductPrice', '.prix_actuel', '.productPrice', '.entry-price', '.cl', '.apparel-price', '.productPricing', '.my-price', '.Price03', '.fichePdtPrix', '.b05', '.now_price', '.actuel', '.product-prices', '.fiche-prix-ttc', '.eur', '.t28', '.product-price-new', '.pvp', '.PriceContainer', '.productprice', '.product-detail-price', '.starting_price', '.price_value', '.art_prix', '.uc-price', '.pb_price', '.priceLarge', '.px_boutique', '.shoe-price', '.fll', '.product__price', '.buySome_dollars', '.detailprix', '.prodPrice', '.TXCD', '.infoBoxContents', '.price-prix', '.price-value', '.fiche_prix_vente', '.salesprice', '.exact-price', '.pricereg', '.productSalePrice', '.a-color-price', '.price_currency', '.big-price', '.cs_product_detail_desc_price', '.prix_produit', '.article-price', '.node-prix'],
        /** ne plus toucher apres ;) **/
        iframe,
        elems = ['head', 'body', 'title', 'base'],
        neeed = {
            desc: document.querySelector('meta[name=description]'),
            path: window.location.protocol + '//' + window.location.host,
            url: window.location.href
        },
        imgs = document.querySelectorAll('img'),
        _link = document.createElement('link'),
        openIframe = function() {
            var l = parseInt(neeed.body.style.left);
            if (l > -iframeWidth) {
                neeed.body.style.left = (l - 10) + 'px';
                iframe.style.right = -l - 200 + 'px'
                setTimeout(function() {
                    openIframe();
                }, 2);
            }
        },
        closeIframe = function() {
            var l = parseInt(neeed.body.style.left);
            if (l < 0) {
                neeed.body.style.left = (l + 10) + 'px';
                iframe.style.right = (-l - 200) + 'px'
                setTimeout(function() {
                    closeIframe();
                }, 2);
            } else {
                var deletes = document.querySelectorAll('.neeed-iframe, .neeed-opacity');
                if (deletes.length)
                    for (var i = 0; i < deletes.length; i++)
                        neeed.body.removeChild(deletes[i]);
                var toremove = document.querySelectorAll('.neeed-focus');
                for (var i = 0; i < toremove.length; i++)
                    toremove[i].parentNode.removeChild(toremove[i]);
            }
        },
        makeIframe = function(data) {
            iframe = document.createElement('iframe');
            iframe.setAttribute('src', url + 'iframe.php?' + data.join('&'));
            iframe.setAttribute('class', 'neeed-iframe');
            iframe.setAttribute('className', 'neeed-iframe');
            iframe.setAttribute('style', 'border: 0; width: ' + iframeWidth + 'px !important; height: ' + window.innerHeight + 'px !important; position: fixed; top: 0; right: -' + iframeWidth + 'px; z-index: 99999999;');
            neeed.body.appendChild(iframe);
            var opacity = document.createElement('div');
            opacity.setAttribute('class', 'neeed-opacity');
            opacity.setAttribute('className', 'neeed-opacity');
            opacity.setAttribute('style', 'position: absolute; top: 0; left: 0; right: 0; bottom: 0; z-index: 99999;');
            neeed.body.appendChild(opacity);
            opacity.onclick = function() {
                closeIframe();
                return false;
            };
            neeed.body.style.position = 'relative';
            neeed.body.style.left = '0';
            window.scroll(0, 0);
            openIframe();
        },
        lavash = function(img, dims, is_img) {
            var wrap = document.createElement('a');
            wrap.setAttribute('class', 'neeed-focus');
            wrap.setAttribute('className', 'neeed-focus');
            wrap.setAttribute('style', 'z-index:99999999; position:absolute; left:' + (dims.left - borderSize) + 'px; top:' + (window.scrollY + dims.top - borderSize) + 'px; width:' + dims.width + 'px; height:' + dims.height + 'px;');
            wrap.setAttribute('href', 'javascript:void;');
            wrap.innerHTML = '<img src="' + url + '../core/img/logo.png" /><br />Ajouter Ã  Neeed';
            wrap.onclick = function() {
                var datas = {},
                    found = false,
                    node = img.parentNode;
                datas.img = is_img ? img.src : img.style.backgroundImage;
                datas.base = neeed.base;
                datas.path = neeed.path;
                datas.desc = neeed.desc;
                datas.title = neeed.title;
                datas.url = neeed.url;
                while (!found && node) {
                    for (var _p in _prices)
                        if (typeof _prices[_p] === 'string') {
                            var price = node.querySelectorAll(_prices[_p]);
                            if (price.length)
                                for (var i = 0; i < price.length; i++) {
                                    var p = price[i].textContent || price[i].innerText || price[i].getAttribute('content');
                                    if (!found && p) {
                                        p = p.trim();
                                        if (p.match(/[0-9]+/)) {
                                            datas.price = p;
                                            found = true;
                                        }
                                    }
                                }
                        }
                    node = node.parentNode;
                }
                var data = [];
                data.push('api_key=xiuhr897Cue');
                for (var i in datas)
                    if (datas[i] !== null)
                        data.push(i + '=' + encodeURIComponent(datas[i]));
                makeIframe(data);
                return false;
            };
            neeed.body.appendChild(wrap);
        };
    for (var i in elems)
        if (typeof elems[i] === "string")
            neeed[elems[i]] = document.querySelector(elems[i]);
    if (neeed.title !== null)
        neeed.title = neeed.title.text;
    if (neeed.desc !== null)
        neeed.desc = neeed.desc.getAttribute('content');
    if (neeed.base !== null)
        neeed.base = neeed.base.getAttribute('href');
    _link.setAttribute('href', url + 'neeed.css');
    _link.setAttribute('rel', 'stylesheet');
    neeed.head.appendChild(_link);
    var found = false;
    for (var i = 0; i < imgs.length; i++) {
        var dims = imgs[i].getBoundingClientRect();
        if (dims.width >= imageWidth && dims.height >= imageHeigth && dims.width <= imageMaxWidth && dims.height <= imageMaxHeigth) {
            lavash(imgs[i], dims, true);
            found = true;
        }
    }
    if (!imgs.length || !found) {
        imgs = document.querySelectorAll('div,section,a');
        for (var i = 0; i < imgs.length; i++) {
            if (imgs[i].style.backgroundImage != '') {
                var dims = imgs[i].getBoundingClientRect();
                if (dims.width >= imageWidth && dims.height >= imageHeigth && dims.width <= imageMaxWidth && dims.height <= imageMaxHeigth) {
                    lavash(imgs[i], dims, false);
                    found = true;
                }
            }
        }
    }
    if (!found) makeIframe(['404', 'neeed_website=' + neeed.path]);
    window.addEventListener('message', function(event) {
        if (event.origin !== "https://neeed.com")
            return;
        closeIframe();
    }, false);

})();