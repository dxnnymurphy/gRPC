from elasticsearch import Elasticsearch
import pandas as pd


class ESClient:
    def __init__(self, ip):
        #change this from hardcoding
        self.es = Elasticsearch(ip, http_auth=("dmurphy", "Queensmead11"), use_ssl=True, verify_certs=False)

    def parseResponse(self,response):
        parsed = []
        rowcounter = 0
        for row in response['rows']:
            parsed.append({})
            columncounter = 0
            for column in response['columns']:
                columnName = response['columns'][columncounter]['name']
                parsed[rowcounter][columnName] = response['rows'][rowcounter][columncounter]
                columncounter += 1
            rowcounter += 1
        return parsed

    def Query(self, query):
        res = self.es.sql.query(body={'query': query, 'fetch_size':10000})
        df=pd.DataFrame(self.parseResponse(res))
        while res.get('cursor'):
            res_new = self.es.sql.query(body={"cursor": res['cursor']})
            res_new['columns'] = res['columns']
            df_new = pd.DataFrame(self.parseResponse(res_new))
            df = pd.concat([df, df_new])
            res = res_new
        return df

class ADUtils:
    def __init__(self):
        pass

    def calculateExecutionVolume(self, df):
        #is this condition correct??
        df2 = df.loc[(df['spec.fix.EXECTYPE'] == 'cancel') | (df['spec.fix.EXECTYPE'] == 'eliminate') | (df['spec.fix.EXECTYPE'] == 'reject')].copy()
        df2['spec.fix.38'] = df2['spec.fix.38'].astype(float)
        df2['spec.fix.44'] = df2['spec.fix.44'].astype(float)
        df2['exec_volume'] = (df2['spec.fix.44'] * df2['spec.fix.38'])
        df_tidy = df2['exec_volume'].dropna()
        X_values = df_tidy.values.reshape(-1,1)
        return X_values
    
    def getAnomalies(self, df, results):
        df['anomaly'] = results
        df_anomalies = df.loc[df['anomaly'] == -1].drop('anomaly')
        return df_anomalies