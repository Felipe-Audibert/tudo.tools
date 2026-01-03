$(document).ready(function() {
    function filterOptions(search, parent) {
        if (search.trim() === '') {
            $(parent).find('.select-search__options li').show();
            return;
        }

        $(parent).find('.select-search__options li').hide();
        $(parent).find('.select-search__options li').filter(function () {
            return $(this).data('label').toLowerCase().includes(search.toLowerCase());
        }).show();
    }

    $(document).on('focus', '.select-search', function() {
        $(this).find('.select-search__options').removeClass('hidden');

        const search = $(this).find('.select-search__input').val();
        filterOptions(search, this);
    });

    $(document).on('blur', '.select-search', function() {
        $(this).find('.select-search__options').addClass('hidden');
    });

    $(document).on('keyup', '.select-search__input', function() {
        const search = $(this).val();
        const parent = $(this).parent('.select-search');

        filterOptions(search, parent);
    });

    $(document).on('click', '.select-search__options li', function() {
        const label = $(this).data('label');
        const value = $(this).data('value');
        const parent = $(this).closest('.select-search');

        $(parent).find('.select-search__input').attr('placeholder', label).val('');
        $(parent).find('.select-search__value').val(value);
        $(parent).find('.select-search__value')[0].dispatchEvent(new Event('change'));

        $(parent).trigger('blur');
    });

    $(document).on('click', '.sidebar-list .summary', function() {
        const item = $(this).parent('li');
        const details = $(this).siblings('.sidebar-list-children');
        const detailsParents = details.parents('.sidebar-list-children');
        if (!details || details.length === 0) {
            return;
        }

        if (item.hasClass('opened')) {
            item.removeClass('opened');
            details.css('maxHeight', 0);
        } else {
            item.addClass('opened');
            details.css('maxHeight', details[0].scrollHeight + 'px');
            detailsParents.css('maxHeight', 'unset');
        }
    })
})