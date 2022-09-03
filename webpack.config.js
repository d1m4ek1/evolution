/* eslint-disable */
const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const FaviconsWebpackPlugin = require("favicons-webpack-plugin");
const copyPlugin = require("copy-webpack-plugin");
const { VueLoaderPlugin } = require("vue-loader");
const Webpack = require("webpack");

const pathEntryFiles = "/src/entry_points/";

module.exports = {
  // devtool: "inline-source-map",
  entry: {
    _index: `${pathEntryFiles}index.js`,
    _autoriz: `${pathEntryFiles}autoriz.js`,
    _profile: `${pathEntryFiles}profile.js`,
    _directory: `${pathEntryFiles}directory.js`,
    _settings: `${pathEntryFiles}settings.js`,
    _messenger: `${pathEntryFiles}messenger.js`,
    _search: `${pathEntryFiles}search.js`,
    _subscriptions: `${pathEntryFiles}subscriptions.js`,
  },
  output: {
    filename: "assets/js/[name].[chunkhash].js",
    publicPath: "/ui/",
    library: "[name]",
    path: path.resolve(__dirname, "ui/"),
    assetModuleFilename: "assets/[hash][ext][query]",
    clean: true,
  },
  mode: "development",
  devServer: {
    static: "/ui/",
    open: true,
    hot: true,
    port: 8080,
  },
  module: {
    rules: [
      {
        test: /\.(sass|css)$/i,
        exclude: /node_modules/,
        use: ["style-loader", "vue-style-loader", "css-loader"],
      },
      {
        test: /\.ts?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
        },
      },
      {
        test: /\.vue$/i,
        loader: "vue-loader",
        options: {
          loader: {
            css: "vue-style-loader!css-loader",
          },
          hotReload: true,
          appendTsSuffixTo: [/\.vue$/],
        },
      },
      {
        test: /\.(html)$/,
        use: ["html-loader"],
      },
      {
        test: /\.(png|svg|jpg|jpeg|gif|ico)$/i,
        type: "asset/resource",
      },
    ],
  },
  plugins: [
    new Webpack.DefinePlugin({
      __VUE_OPTIONS_API__: true,
      __VUE_PROD_DEVTOOLS__: true,
    }),
    new HtmlWebpackPlugin({
      filename: "html/index.html",
      template: "src/index.html",
      chunks: ["_index"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Authoriz.html",
      template: "src/pages/Authorization/Authoriz.html",
      chunks: ["_autoriz"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Profile.html",
      template: "src/pages/Profile/Profile.html",
      chunks: ["_profile"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Directory.html",
      template: "src/pages/Directory/Directory.html",
      chunks: ["_directory"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Settings.html",
      template: "src/pages/Settings/Settings.html",
      chunks: ["_settings"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Messanger.html",
      template: "src/pages/Messenger/Messenger.html",
      chunks: ["_messenger"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Search.html",
      template: "src/pages/Search/Search.html",
      chunks: ["_search"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Subscriptions.html",
      template: "src/pages/Subscriptions/Subscriptions.html",
      chunks: ["_subscriptions"],
    }),
    new copyPlugin({
      patterns: [
        { from: "src/assets/images/", to: "assets/images/" },
        { from: "src/templates/", to: "templates/" },
      ],
    }),
    new FaviconsWebpackPlugin({
      logo: "src/assets/images/main_logo.png",
      prefix: "assets/images/",
    }),
    new VueLoaderPlugin(),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src/"),
      vue: "vue/dist/vue.esm-bundler.js",
    },
    extensions: [".tsx", ".ts", ".js", ".vue"],
  },
};
