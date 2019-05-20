import Taro, { Component, Config } from '@tarojs/taro'
import { View } from '@tarojs/components'
import { AtList, AtListItem } from "taro-ui"

import './index.scss'

class Index extends Component {
  config: Config = {
    navigationBarTitleText: '首页'
  }

  state = {
    topics: [
      {
        id: 1,
        title: '测试1',
        body: 'larabbs 测试内容1'
      }, {
        id: 2,
        title: '测试2',
        body: 'larabbs 测试内容2'
      }, {
        id: 3,
        title: '测试3',
        body: 'larabbs 测试内容3'
      }, {
        id: 4,
        title: '测试4',
        body: 'larabbs 测试内容4'
      }
    ]
  }

  render () {
    return (
      <View>
        <View className='at-article__h3'>
         话题列表
        </View>
        <AtList>
          {
            this.state.topics.map((t) => {
              return (
                <AtListItem
                  key={t.id}
                  title={t.title}
                  note={t.body}
                  thumb='https://iocaffcdn.phphub.org/uploads/avatars/3995_1516760409.jpg?imageView2/1/w/200/h/200'
                />
              )
            })
          }
        </AtList>
      </View>
    )
  }
}

export default Index
