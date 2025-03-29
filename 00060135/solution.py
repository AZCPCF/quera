from bs4 import BeautifulSoup

def process(path):
    with open(path, 'r', encoding='utf-8') as file:
        html_content = file.read()

    soup = BeautifulSoup(html_content, 'html.parser')
    links = soup.find_all('a')  
    return len(links)  
