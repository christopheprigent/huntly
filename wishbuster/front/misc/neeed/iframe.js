// https://neeed.com/bm/iframe.js
(function($) {

    $(document).ready(function() {

        $('#js-cats').on('change', function() {
            $('#js-subcats').hide();
            if ($('option:selected', $(this)).is('[data-genre]'))
                $('#js-subcats').show();
            return false;
        });

        $('#js-form').on('submit', function() {
            if (
                $('[name=neeed_price]').val() != '' &&
                $('[name=neeed_category]').val() != '' &&
                (
                    (
                        $('[name=neeed_category] option:selected').data('genre') == 1 &&
                        $('[name=neeed_sub_category]').val() != ''
                    ) ||
                    typeof $('[name=neeed_category] option:selected').data('genre') == 'undefined'
                )
            )
                return true;
            else
                alert('Merci de remplir tous les champs');
            return false;
        });

    });

})(window.jQuery);