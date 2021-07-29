function GetQueryVars() {
    var vars = {};
    j = 0;
    let q = decodeURIComponent(window.location.search);
    if (q[0] == "?") {
        q = q.slice(1)
        q = "?j=&"+q
    }

    const query = /\?|\&([^=]+)\=([^&]+)/;

    var m = q.match(query);
    for (; m; m = q.match(query)) {
        var l = m[0].toString();
        var len = 0;
        for (len = 0, v = l[len]; v; len++, v = l[len]){}

        q = q.slice(len);

        vars[m[1]] = decodeURIComponent(m[2]).replaceAll('+', ' ');
    }

    return vars;
}

var data = GetQueryVars();
var styles = "/static/style.css"

var html = `
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <title>${data['name']}</title>
    <link rel="stylesheet" href="/static/root.css" type="text/css" media="all">
    <link rel="stylesheet" href="${styles}" type="text/css" media="all">
    <style type="text/css" media="screen">
html, body {
height: 99%;
width: 99%;
background-size: cover;
background-image: url("${data['back']}")
}

@media print {
footer {
    display: none !important;
}
}
    </style>
</head>
<body>
    <section>

        <h1>${data['name']}</h1>
        <img class="logo ${ ! data['logo'] ?'hidden': '' }" style="max-width: content-box;max-height: 200px" src="${data['logo']}"/>
        <h3 class="${!data['slogan']?'hidden':'' }">${data['slogan']}</h3>

        <div id='desc' class="${!data['description']?'hidden':''}">
            <p>${data['description']}</p>
        </div>

        <div>
            <p class="${!data['instagram']?'hidden':''}">
                <a href="${data['instagram_link']}" target="_blank">
                <img class="${ ! data['instagram'] ?'hidden': '' }" src="https://img.icons8.com/fluent/48/000000/instagram-new.png"/> <span>${data['instagram']}</span>
                </a>
            </p>
            <p class="${ ! data['facebook'] ?'hidden': '' }">
                <a href="${data['facebook_link']}" target="_blank">
                    <img src="https://img.icons8.com/fluent/48/000000/facebook-new.png"/> <span>${data['facebook']}</span>
                </a>
            </p>
            <p class="${ ! data['whatsapp'] ?'hidden': '' }">
                <a href="${data['whatsapp_link']}" target="_blank">
                    <img src="https://img.icons8.com/officel/48/000000/whatsapp.png"/> <span>${data['whatsapp']}</span>
                </a>
            </p>
            <p class="${ ! data['email'] ?'hidden': '' }">
                <a href="${data['email_link']}" arget="_blank">
                    <img src="https://img.icons8.com/color/48/000000/send-mass-email.png"/> <span>${data['email']}</span>
                </a>
            </p>
        </div>

    </section>
    <script charset="utf-8" src="/util.js"></script>

    <footer>
        <h3>CopyRight &copy; Yaks Souza</h3>
        <p>&lt;pl4nk1ton@gmail.com&gt;</p>
    </footer>
</body>
`;

document.getElementsByTagName("html")[0].innerHTML = html;
