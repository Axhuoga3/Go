import random
from datetime import datetime, timedelta

# 常见姓氏和名字数据集（可根据需要扩展）
last_names = [
    '王', '李', '张', '刘', '陈', '杨', '黄', '赵', '周', '吴',
    '徐', '孙', '马', '朱', '胡', '林', '郭', '何', '高', '罗'
]

first_names_male = [
    '伟', '强', '磊', '军', '勇', '杰', '涛', '斌', '超', '鹏',
    '浩', '亮', '刚', '健', '明', '平', '辉', '建国', '建军', '志强'
]

first_names_female = [
    '芳', '娜', '敏', '静', '丽', '娟', '艳', '玲', '霞', '婷',
    '雪', '慧', '玉', '英', '红', '梅', '倩', '欣', '淑华', '秀兰'
]

# 地址相关数据
provinces = ['北京市', '上海市', '广东省', '浙江省', '江苏省']
cities = {
    '北京市': ['市辖区'],
    '上海市': ['市辖区'],
    '广东省': ['广州市', '深圳市', '东莞市'],
    '浙江省': ['杭州市', '宁波市', '温州市'],
    '江苏省': ['南京市', '苏州市', '无锡市']
}

streets = [
    '中山路', '解放路', '人民路', '建设路', '和平街',
    '新华街', '朝阳路', '东风路', '滨海大道', '创业路'
]

def generate_gender():
    return random.choice(['男', '女'])

def generate_birthdate(min_age=18, max_age=80):
    now = datetime.now()
    start_date = now - timedelta(days=max_age*365)
    end_date = now - timedelta(days=min_age*365)
    random_date = start_date + (end_date - start_date) * random.random()
    return random_date.strftime("%Y-%m-%d")

def generate_name(gender):
    last_name = random.choice(last_names)
    if gender == '男':
        first_name = random.choice(first_names_male)
    else:
        first_name = random.choice(first_names_female)
    return last_name + first_name

def generate_address():
    province = random.choice(provinces)
    city = random.choice(cities[province])
    street = random.choice(streets)
    number = str(random.randint(1, 999)) + '号'
    return f"{province}{city}{street}{number}"

def generate_phone_number():
    prefix = random.choice(['13', '15', '18', '19'])
    suffix = ''.join(random.choices('0123456789', k=9))
    return prefix + suffix

def generate_email(name):
    domains = ['example.com', 'testmail.com', 'dummy.org']
    name_part = name.replace(' ', '').lower()
    return f"{name_part}{random.randint(10,99)}@{random.choice(domains)}"

def generate_virtual_identity():
    gender = generate_gender()
    name = generate_name(gender)
    return {
        "姓名": name,
        "性别": gender,
        "出生日期": generate_birthdate(),
        "地址": generate_address(),
        "手机号码": generate_phone_number(),
        "电子邮箱": generate_email(name)
    }

if __name__ == "__main__":
    identity = generate_virtual_identity()
    for key, value in identity.items():
        print(f"{key}: {value}")
