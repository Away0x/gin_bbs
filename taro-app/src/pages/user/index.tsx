import Taro, { Component, Config } from '@tarojs/taro'
import { View } from '@tarojs/components'

import './index.scss'

class User extends Component {
  config: Config = {
    navigationBarTitleText: '我的'
  }

  render() {
    return (
      <View className="container">
        我的页面
      </View>
    )
  }
}

export default User
