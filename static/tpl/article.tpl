<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{header}}</title>
    <link href="../css/app.css" rel="stylesheet">
    <link href="../css/customer.css" rel="stylesheet">
    <link href="../css/font-awesome.min.css" rel="stylesheet">
    <script src="../js/jquery-1.9.1.min.js"></script>
    <script src="../js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">
    <div class="dist-main-container">
        <div class="col-md-12">
            <div class="dist-article-content">
                <div class="title-article">
                    <h1>{{title}}</h1>
                </div>
                <div class="tag-article">
                    <span class="label"><i class="fa fa-tags"></i> {{created_at}}</span>
                    <span class="label"><i class="fa fa-user"></i> {{author}}</span>
                    <span class="label"><i class="fa fa-eye"></i> {{views}}</span>
                </div>
                <div id="content">
                    <textarea style="display:none;">{{body}}</textarea>
                </div>
            </div>
        </div>
    </div>
</div>
<link rel="stylesheet" href="../css/editormd.css" />
<script src="../js/lib/marked.min.js"></script>
<script src="../js/lib/prettify.min.js"></script>
<script src="../js/lib/raphael.min.js"></script>
<script src="../js/lib/underscore.min.js"></script>
<script src="../js/lib/sequence-diagram.min.js"></script>
<script src="../js/lib/flowchart.min.js"></script>
<script src="../js/lib/jquery.flowchart.min.js"></script>
<script src="../js/editormd.min.js"></script>
<script type="text/javascript">
    $(function() {
        editormd.markdownToHTML("content", {
            htmlDecode      : "style,script,iframe",
            emoji           : true,
            taskList        : true,
            tex             : true,
            flowChart       : true,
            sequenceDiagram : true,
        });
    })
</script>
</body>
</html>