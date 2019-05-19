import '@tarojs/async-await'
import Taro, { Component, Config } from '@tarojs/taro'
import { Provider } from '@tarojs/redux'

import Index from './pages/index'

import configStore from './store'

import 'taro-ui/dist/style/index.scss'
import './app.scss'

const store = configStore()

class App extends Component {
  config: Config = {
    pages: [
      'pages/index/index',
      'pages/user/index'
    ],
    window: {
      backgroundTextStyle: 'light',
      navigationBarBackgroundColor: '#fff',
      navigationBarTitleText: 'WeChat',
      navigationBarTextStyle: 'black'
    },
    tabBar: {
      color: '#707070',
      selectedColor: '#00b5ad',
      list: [
        {
          pagePath: 'pages/index/index',
          text: '首页',
          iconPath: 'assets/images/index.png',
          selectedIconPath: 'assets/images/index_selected.png'
        },
        {
          pagePath: 'pages/user/index',
          text: '我的',
          iconPath: 'assets/images/user.png',
          selectedIconPath: 'assets/images/user_selected.png'
        }
      ]
    }
  }

  async componentDidMount () {
    const resource = await Taro.login();
    console.log(resource)
  }

  componentDidShow () {}

  componentDidHide () {}

  componentDidCatchError () {}

  render () {
    return (
      <Provider store={store}>
        <Index />
      </Provider>
    )
  }
}

Taro.render(<App />, document.getElementById('app'))
