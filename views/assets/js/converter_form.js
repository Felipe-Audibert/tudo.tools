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

    const selectSearch = $('.select-search');

    $(selectSearch).on('focusin',function () {
        $(this).find('.select-search__options').removeClass('hidden');

        const search = $(this).find('.select-search__input').val();
        filterOptions(search, this);
    });

    $(selectSearch).on('focusout', function () {
        $(this).find('.select-search__options').addClass('hidden');
    });

    $('.select-search__input').on('keyup', function () {
        const search = $(this).val();
        const parent = $(this).parent('.select-search');

        filterOptions(search, parent);
    });

    $('.select-search__options li').on('click', function () {
        const label = $(this).data('label');
        const value = $(this).data('value');
        const parent = $(this).closest('.select-search');

        $(parent).find('.select-search__input').attr('placeholder', label).val('');
        $(parent).find('.select-search__value').val(value);
        $(parent).find('.select-search__value')[0].dispatchEvent(new Event('change'));

        $(parent).trigger('focusout');
    });

    $('.select-search__value').on('change', function() {
        const parent = $(this).closest('.select-search');
        const label = $(parent).find(`li[data-value="${$(this).val()}"]`).data('label');
        $(parent).find('.select-search__input').attr('placeholder', label);
    });

    $('.switch-button').on('click', function() {
        const form = $(this).closest('form');
        const fromInput = $(form).find('#from-input');
        const toOutput = $(form).find('#to-output');
        const fromSelect = $(form).find('#from-select');
        const toSelect = $(form).find('#to-select');

        toOutput.html() !== '' && fromInput.val(toOutput.html());
        toOutput.html('');

        const fromSelectVal = fromSelect.val();
        fromSelect.val(toSelect.val());
        toSelect.val(fromSelectVal);


        form.find('.select-search__value').trigger('change');
        $(form).find('.select-search__value')[0].dispatchEvent(new Event('change'));
    });
});