(function($) {
    function addLists(list) {
        if (list.pub) {
            opt = document.createElement('optgroup');
            opt.setAttribute('label', "Listes publiques");

            for (var idx in list.pub) {
                var publist = list.pub[idx];
                o = document.createElement('option');
                o.innerText = publist;
                opt.append(o);
            }
            $("#hunt_list").append(opt);
        }
        if (list.priv) {
            opt2 = document.createElement('optgroup');
            opt2.setAttribute('label', "Listes Privèes");

            for (var idx2 in list.priv) {
                var publist2 = list.priv[idx2];
                o = document.createElement('option');
                o.innerText = publist2;
                opt2.append(o);
            }
            $("#hunt_list").append(opt2);
        }


    }
    $(document).ready(function() {
        console.log(window.HuntProduct);
        // TODO : REMOVE MOCK
        usr_lst = {
            "pub": ["toto", "tutu"],
            "priv": ["priv-toto", "priv-tutu"]
        };

        var match,
            pl = /\+/g, // Regex for replacing addition symbol with a space
            search = /([^&=]+)=?([^&]*)/g,
            decode = function(s) { return decodeURIComponent(s.replace(pl, " ")); },
            query = window.location.search.substring(1);

        urlParams = {};
        while (match = search.exec(query))
            urlParams[decode(match[1])] = decode(match[2]);

        product_label = "";
        product_ref = "product_ref";
        product_img = urlParams.img;
        product_price = urlParams.price || "";
        product_price = product_price.replace(/\,/, '.').replace(/[^\d,\.]/gi, '');
        console.log("From website :");
        console.log(urlParams);

        //pre-fill with custom data
        addLists(usr_lst);
        $("#hunt_img").attr("src", product_img);
        $("#hunt-price-input").val(product_price);

        $('#js-cats').on('change', function() {
            $('#js-subcats').hide();
            if ($('option:selected', $(this)).is('[data-genre]'))
                $('#js-subcats').show();
            return false;
        });

        $('#js-form').on('submit', function() {
            // if (
            //     $('[name=hunt_price]').val() != '' &&
            //     $('[name=hunt_category]').val() != '' &&
            //     (
            //         (
            //             $('[name=hunt_category] option:selected').data('genre') == 1 &&
            //             $('[name=hunt_sub_category]').val() != ''
            //         ) ||
            //         typeof $('[name=hunt_category] option:selected').data('genre') == 'undefined'
            //     )
            // ) {
            // TODO CALL BACK
            alert("TODO CALL BACK");
            return false;
            // } else
            //     alert('Merci de remplir tous les champs');
            // return false;
        });

    });

})(window.jQuery);