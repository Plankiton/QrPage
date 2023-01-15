function load(data) {
  var styles = "style.css"
  if (data['style'])
    styles = data['style'];

  var html = `
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <title>${data['name']}</title>
    <link rel="stylesheet" href="/css/root.css" type="text/css" media="all">
    <link rel="stylesheet" href="/css/${styles}" type="text/css" media="all">
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
        <img class="logo ${!data['logo'] ? 'hidden' : ''}" style="max-width: content-box;max-height: 200px" src="${data['logo']}"/>
        <h3 class="${!data['slogan'] ? 'hidden' : ''}">${data['slogan']}</h3>

        <div id='desc' class="${!data['description'] ? 'hidden' : ''}">
            <p>${data['description']}</p>
        </div>

        <div>
            <p class="${!data['instagram'] ? 'hidden' : ''}">
                <a href="${data['instagram_link']}" target="_blank">
                <img class="${!data['instagram'] ? 'hidden' : ''}" src="https://img.icons8.com/fluent/48/000000/instagram-new.png"/> <span>${data['instagram']}</span>
                </a>
            </p>
            <p class="${!data['facebook'] ? 'hidden' : ''}">
                <a href="${data['facebook_link']}" target="_blank">
                    <img src="https://img.icons8.com/fluent/48/000000/facebook-new.png"/> <span>${data['facebook']}</span>
                </a>
            </p>
            <p class="${!data['whatsapp'] ? 'hidden' : ''}">
                <a href="${data['whatsapp_link']}" target="_blank">
                    <img src="https://img.icons8.com/officel/48/000000/whatsapp.png"/> <span>${data['whatsapp']}</span>
                </a>
            </p>
            <p class="${!data['email'] ? 'hidden' : ''}">
                <a href="${data['email_link']}" arget="_blank">
                    <img src="https://img.icons8.com/color/48/000000/send-mass-email.png"/> <span>${data['email']}</span>
                </a>
            </p>
        </div>

    </section>

    <script>
      FIRST_RUN = false;
    </script>
    <script charset="utf-8" src="/js/qrpage.js"></script>

    <footer>
        <h3>CopyRight &copy; Yaks Souza</h3>
        <p>&lt;yaks.souza@gmail.com&gt;</p>
    </footer>
</body>
`;

  document.getElementsByTagName("html")[0].innerHTML = html;
}

function GetQueryVars() {
  var vars = {};
  j = 0;
  let q = decodeURIComponent(window.location.search);
  if (q[0] == "?") {
    q = q.slice(1)
    q = "?j=&" + q
  }

  const query = /\?|\&([^=]+)\=([^&]+)/;

  var m = q.match(query)
  for (; m; m = q.match(query)) {
    var l = m[0].toString();
    var len = 0;
    for (len = 0, v = l[len]; v; len++, v = l[len]) { }

    q = q.slice(len);

    vars[m[1]] = decodeURIComponent(m[2]).replaceAll('+', ' ');
  }

  return vars;
}

function convertToServerData(d) {
  return {
    business: {
      name: d.name,
      description: d.description,
      corp_phrase: d.slogan,
      logo_image: d.logo
    },
    content: {
      background_image: d.back,
      style: d.style
    },
    social: [
      {
        type: "instagram",
        title: d.instagram,
        back_link: d.instagram_link,
      },
      {
        type: "email",
        title: d.email,
        back_link: d.email_link,
      },
      {
        type: "whatsapp",
        title: d.whatsapp,
        back_link: d.whatsapp_link,
      },
      {
        type: "facebook",
        title: d.facebook,
        back_link: d.facebook_link,
      },
    ]
  }
}


function convertToData(d) {
  var instagram = null;
  var instagram_link = null;
  var facebook = null;
  var facebook_link = null;
  var whatsapp = null;
  var whatsapp_link = null;
  var email = null;
  var email_link = null;

  for (let i = 0; i < d.social.length; i++) {
    if (d.social[i].type == "instagram") {
      instagram = d.social[i].title;
      instagram_link = d.social[i].back_link;
    }

    if (d.social[i].type == "facebook") {
      facebook = d.social[i].title;
      facebook_link = d.social[i].back_link;
    }

    if (d.social[i].type == "whatsapp") {
      whatsapp = d.social[i].title;
      whatsapp_link = d.social[i].back_link;
    }

    if (d.social[i].type == "email") {
      email = d.social[i].title;
      email_link = d.social[i].back_link;
    }
  }

  return {
    back: d.content.background_image,
    logo: d.business.logo_image,
    slogan: d.business.corp_phrase,
    description: d.business.description,
    name: d.business.name,

    instagram: instagram,
    instagram_link: instagram_link,

    email: email,
    email_link: email_link,

    facebook: facebook,
    facebook_link: facebook_link,

    whatsapp: whatsapp,
    whatsapp_link: whatsapp_link,
  }
}

var data = GetQueryVars();

function getDataByLink(slugCode) {
  axios.get(`/api/i/${slugCode}`).
    then(r => {
      server_data = r.data.data;
      console.log("server data", server_data);
      data = convertToData(server_data)
      console.log("conv data", data);

      load(data);
    }).
    catch(e => {
      console.log(e)
    })
}

console.log(convertToServerData(data));

console.log(data);
if (data['in'] != undefined)
  getDataByLink(data.in);
else if (FIRST_RUN)
  load(data);
