$.ajax({


    headers: {
        'Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRFU1QiLCJpZGVudGl0eSI6IueUqOaItyIsImV4cCI6MTY0NTQyNzU3MCwiaXNzIjoiZG91QmFuIiwibmJmIjoxNjQ1NDIwMzEwLCJzdWIiOiJIYXJkZXdzIn0.mjzy71U65Y3XryrsxgOXkzlTD7wrcCPwmaCHGHL74Pk' //此处放置请求到的用户token
    },

    type: 'post',
    url: ``,
    data: {
        cate_id: 2,
        get_count: 1,
        city_name: '郑州'
    },


})
