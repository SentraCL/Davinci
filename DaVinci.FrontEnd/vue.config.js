// eslint-disable-next-line no-undef
//const configureAPI = require("./src/server/davinci.server");
const path = require("path");

module.exports = {
  chainWebpack: config => {
    config.resolve.alias.set("@", path.join(__dirname, "/vue/"));
  },

  lintOnSave: false,
  /*
  devServer: {
    before: configureAPI
  },
  */
  //Ruta de produccion
  outputDir: "../dist/public_html",

  pages: {
    index: {

      entry: "vue/main.js",
      // the source template

      // output as dist/index.html
      filename: "index.html",
      // chunks to include on this page, by default includes
      // extracted common chunks and vendor chunks.
      chunks: ["chunk-vendors", "chunk-common", "index"]
    },
    project: {
      // entry for the page
      entry: "vue/submain.js",
      // the source template

      // output as dist/index.html
      filename: "project.html",
      // chunks to include on this page, by default includes
      // extracted common chunks and vendor chunks.
      chunks: ["chunk-vendors", "chunk-common", "project"]
    },    
    // when using the entry-only string format,
    // template is inferred to be `public/subpage.html`
    // and falls back to `public/index.html` if not found.
    // Output filename is inferred to be `subpage.html`.
    //subpage: 'src/subpage/main.js'
  }
};
