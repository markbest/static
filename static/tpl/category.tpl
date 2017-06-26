<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>www.markbest.site</title>
    <link rel="stylesheet" type="text/css" href="../css/app.css" />
    <link rel="stylesheet" type="text/css" href="../css/customer.css" />
    <script src="../js/jquery-1.9.1.min.js"></script>
    <script type="text/javascript">
        $(function(){
            $('.tree li:has(ul)').addClass('parent_li').find(' > span').attr('title', 'Collapse this branch');
            $('.tree li.parent_li > span').on('click', function (e) {
                var children = $(this).parent('li.parent_li').find(' > ul > li');
                if (children.is(":visible")) {
                    children.hide('fast');
                    $(this).attr('title', 'Expand this branch').find(' > i').addClass('icon-plus-sign').removeClass('icon-minus-sign');
                } else {
                    children.show('fast');
                    $(this).attr('title', 'Collapse this branch').find(' > i').addClass('icon-minus-sign').removeClass('icon-plus-sign');
                }
                e.stopPropagation();
            });

            $('.category-bar a').click(function(event){
                event.preventDefault();

                $('.category-bar a').removeClass('active');
                $(this).addClass('active');
                var href = $(this).attr('href');
                $('#content').attr('src', href);
            });
        });
    </script>
</head>
<body>
    <div class="category-bar">
        {{category}}
    </div>
    <div class="article-bar">
        <iframe id="content" class="col-md-12" src="{{default_page}}"></iframe>
    </div>
    <div id="footer">
        {{footer}}
    </div>
</body>
</html>

