window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
        {"name":"sprint", "url":"sprint.swagger.json"},
        {"name":"issue", "url":"issue.swagger.json"},
        {"name":"team", "url":"team.swagger.json"},
        {"name":"project", "url":"project.swagger.json"},
        {"name":"department", "url":"department.swagger.json"}
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
