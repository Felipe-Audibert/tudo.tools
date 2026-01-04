$(document).ready(function() {
    $('.sidebar-list .summary').on('click', function() {
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