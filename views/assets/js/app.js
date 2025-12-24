$(document).ready(function() {
    $(document).on('click', '.sidebar-list .summary', function() {
        const item = $(this).parent('li');
        const details = $(this).siblings('.sidebar-list-children');
        const detailsParents = details.parents('.sidebar-list-children');
        if (!details || details.length === 0) {
            return;
        }

        console.log(details.parents('.sidebar-list-children'));

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