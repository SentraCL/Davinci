// eslint-disable-next-line no-undef
const path = require("path");

const { VueLoaderPlugin } = require("vue-loader");

const webpack = require("webpack");

//const { ProgressPlugin } = require("progress");

module.exports = {
  mode: "development",
  context: __dirname,
  devtool: "cheap-module-eval-source-map",
  node: {
    setImmediate: false,
    dgram: "empty",
    fs: "empty",
    net: "empty",
    tls: "empty",
    child_process: "empty",
  },
  output: {
    filename: "[name].js",
    //Ruta de Desarrollo
    publicPath: "/",

    globalObject: "this",
  },
  resolve: {
    alias: {
      "@": __dirname + "/vue/",
      vue$: "vue/dist/vue.esm.js",
    },
    extensions: [".mjs", ".js", ".jsx", ".vue", ".json", ".wasm"],
    modules: [
      "node_modules",
      __dirname + "/node_modules",
      __dirname + "/node_modules/@vue/cli-service/node_modules",
    ],
  },
  resolveLoader: {
    modules: [
      __dirname + "/node_modules/@vue/cli-plugin-babel/node_modules",
      "node_modules",
      __dirname + "/node_modules",
      __dirname + "/node_modules/@vue/cli-service/node_modules",
    ],
  },
  module: {
    noParse: /^(vue|vue-router|vuex|vuex-router-sync)$/,
    rules: [
      /* config.module.rule('vue') */
      {
        test: /\.vue$/,
        use: [
          /* config.module.rule('vue').use('cache-loader') */
          {
            loader: "cache-loader",
            options: {
              cacheDirectory: __dirname + "/node_modules/.cache/vue-loader",
              cacheIdentifier: "32e27042",
            },
          },
          /* config.module.rule('vue').use('vue-loader') */
          {
            loader: "vue-loader",
            options: {
              compilerOptions: {
                preserveWhitespace: false,
              },
              cacheDirectory: __dirname + "/node_modules/.cache/vue-loader",
              cacheIdentifier: "32e27042",
            },
          },
        ],
      },
      /* config.module.rule('images') */
      {
        test: /\.(png|jpe?g|gif|webp)(\?.*)?$/,
        use: [
          /* config.module.rule('images').use('url-loader') */
          {
            loader: "url-loader",
            options: {
              limit: 4096,
              fallback: {
                loader: "file-loader",
                options: {
                  name: "img/[name].[hash:8].[ext]",
                },
              },
            },
          },
        ],
      },

      /* config.module.rule('svg') */
      {
        test: /\.(svg)(\?.*)?$/,
        use: [
          /* config.module.rule('svg').use('file-loader') */
          {
            loader: "file-loader",
            options: {
              name: "img/[name].[hash:8].[ext]",
            },
          },
        ],
      },
      /* config.module.rule('media') */
      {
        test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
        use: [
          /* config.module.rule('media').use('url-loader') */
          {
            loader: "url-loader",
            options: {
              limit: 4096,
              fallback: {
                loader: "file-loader",
                options: {
                  name: "media/[name].[hash:8].[ext]",
                },
              },
            },
          },
        ],
      },
      /* config.module.rule('fonts') */
      {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/i,
        use: [
          /* config.module.rule('fonts').use('url-loader') */
          {
            loader: "url-loader",
            options: {
              limit: 4096,
              fallback: {
                loader: "file-loader",
                options: {
                  name: "fonts/[name].[hash:8].[ext]",
                },
              },
            },
          },
        ],
      },
      /* config.module.rule('pug') */
      {
        test: /\.pug$/,
        oneOf: [
          /* config.module.rule('pug').oneOf('pug-vue') */
          {
            resourceQuery: /vue/,
            use: [
              /* config.module.rule('pug').oneOf('pug-vue').use('pug-plain-loader') */
              {
                loader: "pug-plain-loader",
              },
            ],
          },
          /* config.module.rule('pug').oneOf('pug-template') */
          {
            use: [
              /* config.module.rule('pug').oneOf('pug-template').use('raw') */
              {
                loader: "raw-loader",
              },
              /* config.module.rule('pug').oneOf('pug-template').use('pug-plain') */
              {
                loader: "pug-plain-loader",
              },
            ],
          },
        ],
      },
      /* config.module.rule('css') */
      {
        test: /\.css$/,
        oneOf: [
          /* config.module.rule('css').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('css').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('css').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
            ],
          },
          /* config.module.rule('css').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('css').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('css').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
            ],
          },
          /* config.module.rule('css').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('css').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('css').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
            ],
          },
          /* config.module.rule('css').oneOf('normal') */
          {
            use: [
              /* config.module.rule('css').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('css').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('postcss') */
      {
        test: /\.p(ost)?css$/,
        oneOf: [
          /* config.module.rule('postcss').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('postcss').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('postcss').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
            ],
          },
          /* config.module.rule('postcss').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('postcss').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('postcss').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
            ],
          },
          /* config.module.rule('postcss').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('postcss').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('postcss').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
            ],
          },
          /* config.module.rule('postcss').oneOf('normal') */
          {
            use: [
              /* config.module.rule('postcss').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('postcss').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('scss') */
      {
        test: /\.scss$/,
        oneOf: [
          /* config.module.rule('scss').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('scss').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('scss').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('scss').oneOf('vue-modules').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('scss').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('scss').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('scss').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('scss').oneOf('vue').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('scss').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('scss').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('scss').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('scss').oneOf('normal-modules').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('scss').oneOf('normal') */
          {
            use: [
              /* config.module.rule('scss').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('scss').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('scss').oneOf('normal').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('sass') */
      {
        test: /\.sass$/,
        oneOf: [
          /* config.module.rule('sass').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('sass').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('sass').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('sass').oneOf('vue-modules').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                  indentedSyntax: true,
                },
              },
            ],
          },
          /* config.module.rule('sass').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('sass').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('sass').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('sass').oneOf('vue').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                  indentedSyntax: true,
                },
              },
            ],
          },
          /* config.module.rule('sass').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('sass').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('sass').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('sass').oneOf('normal-modules').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                  indentedSyntax: true,
                },
              },
            ],
          },
          /* config.module.rule('sass').oneOf('normal') */
          {
            use: [
              /* config.module.rule('sass').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('sass').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('sass').oneOf('normal').use('sass-loader') */
              {
                loader: "sass-loader",
                options: {
                  sourceMap: false,
                  indentedSyntax: true,
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('less') */
      {
        test: /\.less$/,
        oneOf: [
          /* config.module.rule('less').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('less').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('less').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('less').oneOf('vue-modules').use('less-loader') */
              {
                loader: "less-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('less').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('less').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('less').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('less').oneOf('vue').use('less-loader') */
              {
                loader: "less-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('less').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('less').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('less').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('less').oneOf('normal-modules').use('less-loader') */
              {
                loader: "less-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
          /* config.module.rule('less').oneOf('normal') */
          {
            use: [
              /* config.module.rule('less').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('less').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('less').oneOf('normal').use('less-loader') */
              {
                loader: "less-loader",
                options: {
                  sourceMap: false,
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('stylus') */
      {
        test: /\.styl(us)?$/,
        oneOf: [
          /* config.module.rule('stylus').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('stylus').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('stylus').oneOf('vue-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('stylus').oneOf('vue-modules').use('stylus-loader') */
              {
                loader: "stylus-loader",
                options: {
                  sourceMap: false,
                  preferPathResolver: "webpack",
                },
              },
            ],
          },
          /* config.module.rule('stylus').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('stylus').oneOf('vue').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('stylus').oneOf('vue').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('stylus').oneOf('vue').use('stylus-loader') */
              {
                loader: "stylus-loader",
                options: {
                  sourceMap: false,
                  preferPathResolver: "webpack",
                },
              },
            ],
          },
          /* config.module.rule('stylus').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('stylus').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('stylus').oneOf('normal-modules').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                  modules: true,
                  localIdentName: "[name]_[local]_[hash:base64:5]",
                },
              },
              /* config.module.rule('stylus').oneOf('normal-modules').use('stylus-loader') */
              {
                loader: "stylus-loader",
                options: {
                  sourceMap: false,
                  preferPathResolver: "webpack",
                },
              },
            ],
          },
          /* config.module.rule('stylus').oneOf('normal') */
          {
            use: [
              /* config.module.rule('stylus').oneOf('normal').use('vue-style-loader') */
              {
                loader: "vue-style-loader",
                options: {
                  sourceMap: false,
                  shadowMode: false,
                },
              },
              /* config.module.rule('stylus').oneOf('normal').use('css-loader') */
              {
                loader: "css-loader",
                options: {
                  sourceMap: false,
                  importLoaders: 1,
                },
              },
              /* config.module.rule('stylus').oneOf('normal').use('stylus-loader') */
              {
                loader: "stylus-loader",
                options: {
                  sourceMap: false,
                  preferPathResolver: "webpack",
                },
              },
            ],
          },
        ],
      },
      /* config.module.rule('js') */
      {
        test: /\.m?jsx?$/,
        exclude: [
          function() {
            /* omitted long function */
          },
        ],
        use: [
          /* config.module.rule('js').use('cache-loader') */
          {
            loader: "cache-loader",
            options: {
              cacheDirectory: __dirname + "/node_modules/.cache/babel-loader",
              cacheIdentifier: "329f2592",
            },
          },
          /* config.module.rule('js').use('babel-loader') */
          {
            loader: "babel-loader",
          },
        ],
      },
    ],
  },
  plugins: [new VueLoaderPlugin()],
  entry: {
    index: [__dirname + "/vue/main.js"],
    project: [__dirname + "/vue/submain.js"],
  },
};
