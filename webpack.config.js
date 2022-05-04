/* eslint-disable */
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const FaviconsWebpackPlugin = require('favicons-webpack-plugin');
const copyPlugin = require('copy-webpack-plugin');
const {
  VueLoaderPlugin,
} = require('vue-loader');

const pathEntryFiles = '/src/entry_points/';

module.exports = {
  devtool: 'inline-source-map',
  entry: {
    _index: `${pathEntryFiles}_index/index.js`,
    _autoriz: `${pathEntryFiles}_autoriz/autoriz.js`,
    _profile: `${pathEntryFiles}_profile/profile.js`,
    _directory: `${pathEntryFiles}_directory/directory.js`,
    _settings: `${pathEntryFiles}_settings/settings.js`,
    _messages: `${pathEntryFiles}__community/_messages/messages.js`,
    _music: `${pathEntryFiles}__community/_music/music.js`,
    _subscriptions: `${pathEntryFiles}__community/_subscriptions/subscriptions.js`,
    _shop: `${pathEntryFiles}__commerce/_shop/shop.js`,
    _orders: `${pathEntryFiles}__commerce/_orders/orders.js`,
  },
  output: {
    filename: 'assets/js/[name].[chunkhash].js',
    publicPath: '/ui/',
    library: '[name]',
    path: path.resolve(__dirname, 'ui/'),
    assetModuleFilename: 'assets/[hash][ext][query]',
    clean: true,
  },
  mode: 'development',
  devServer: {
    static: '/ui/',
    open: true,
    hot: true,
    port: 8080,
  },
  module: {
    rules: [
      {
        test: /\.(sass|css)$/i,
        exclude: /node_modules/,
        use: ['style-loader', 'vue-style-loader', 'css-loader'],
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
      {
        test: /\.vue$/i,
        loader: 'vue-loader',
        options: {
          loader: {
            css: 'vue-style-loader!css-loader',
          },
          hotReload: true,
        },
      },
      {
        test: /\.(html)$/,
        use: ['html-loader'],
      },
      {
        test: /\.(png|svg|jpg|jpeg|gif|ico)$/i,
        type: 'asset/resource',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: 'html/index.html',
      template: 'src/index.html',
      chunks: ['_index'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Autoriz.html',
      template: 'src/pages/Autorization/Autoriz.html',
      chunks: ['_autoriz'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Profile.html',
      template: 'src/pages/Profile/Profile.html',
      chunks: ['_profile'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Directory.html',
      template: 'src/pages/Directory/Directory.html',
      chunks: ['_directory'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Settings.html',
      template: 'src/pages/Settings/Settings.html',
      chunks: ['_settings'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Messages.html',
      template: 'src/pages/__community/Messages/Messages.html',
      chunks: ['_messages'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Music.html',
      template: 'src/pages/__community/Music/Music.html',
      chunks: ['_music'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Subscriptions.html',
      template: 'src/pages/__community/Subscriptions/Subscriptions.html',
      chunks: ['_subscriptions'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Shop.html',
      template: 'src/pages/__commerce/Shop/Shop.html',
      chunks: ['_shop'],
    }),
    new HtmlWebpackPlugin({
      filename: 'html/Orders.html',
      template: 'src/pages/__commerce/Orders/Orders.html',
      chunks: ['_orders'],
    }),
    new copyPlugin({
      patterns: [
        { from: 'src/assets/images/', to: 'assets/images/' },
      ],
    }),
    new FaviconsWebpackPlugin({
      logo: 'src/assets/images/main_logo.png',
      prefix: 'assets/images/',
    }),
    new VueLoaderPlugin(),
  ],
  resolve: {
    alias: {
      vue$: 'vue/dist/vue.js',
    },
  },
};
