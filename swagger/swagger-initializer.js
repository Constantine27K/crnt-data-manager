window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
        {"name":"sprint", "url":"sprint.swagger.json"},
        {"name":"task", "url":"task.swagger.json"},
        {"name":"subtask", "url":"subtask.swagger.json"},
        {"name":"story", "url":"story.swagger.json"},
        {"name":"epic", "url":"epic.swagger.json"}
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
