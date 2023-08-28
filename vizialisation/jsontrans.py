import json

with open('testpy.json', 'r') as json_file:
    json_data = json_file.read()
    data = json.loads(json_data)

node_dict = {}

for item in data:
    n = item['n']
    r = item['r']
    m = item['m']

    if n['elementId'] not in node_dict:
        node_dict[n['elementId']] = {
            'name': f"Node {n['properties']['idA']}",
            'value': n['properties']['idA'],
            'children': []
        }

    if r['type'] == 'DIRECT' and r['startNodeElementId'] == n['elementId']:
        if m['elementId'] not in node_dict:
            node_dict[m['elementId']] = {
                'name': f"Node {m['properties']['idA']}",
                'value': m['properties']['idA'],
                'children': []
            }

        node_dict[n['elementId']]['children'].append(node_dict[m['elementId']])

root_nodes = [node for node in node_dict.values() if 'elementId' not in node]

with open('transformed_data.json', 'w') as output_file:
    json.dump(root_nodes, output_file, indent=2)
