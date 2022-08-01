/* eslint-disable */
const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const FaviconsWebpackPlugin = require("favicons-webpack-plugin");
const copyPlugin = require("copy-webpack-plugin");
const { VueLoaderPlugin } = require("vue-loader");
const Webpack = require("webpack");

const pathEntryFiles = "/src/entry_points/";

module.exports = {
  devtool: "inline-source-map",
  entry: {
    _index: `${pathEntryFiles}_index/index.js`,
    _autoriz: `${pathEntryFiles}_autoriz/autoriz.js`,
    _profile: `${pathEntryFiles}_profile/profile.js`,
    _directory: `${pathEntryFiles}_directory/directory.js`,
    _settings: `${pathEntryFiles}_settings/settings.js`,
    _insocial: `${pathEntryFiles}__projects/insocial.js`,
    _inmusic: `${pathEntryFiles}__projects/inmusic.js`,
    _inbeats: `${pathEntryFiles}__projects/inbeats.js`,
    _subscriptions: `${pathEntryFiles}__commerce/subscriptions.js`,
    _orders: `${pathEntryFiles}__commerce/orders.js`,
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
      filename: "html/inSocial.html",
      template: "src/pages/__projects/inSocial/inSocial.html",
      chunks: ["_insocial"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/inMusic.html",
      template: "src/pages/__projects/inMusic/inMusic.html",
      chunks: ["_inmusic"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/inBeats.html",
      template: "src/pages/__projects/inBeats/inBeats.html",
      chunks: ["_inbeats"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Subscriptions.html",
      template: "src/pages/__commerce/Subscriptions/Subscriptions.html",
      chunks: ["_subscriptions"],
    }),
    new HtmlWebpackPlugin({
      filename: "html/Orders.html",
      template: "src/pages/__commerce/Orders/Orders.html",
      chunks: ["_orders"],
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
