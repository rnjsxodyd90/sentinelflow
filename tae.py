import pandas as pd

def clean_data(file_path):
    print(f"reading {file_path}")

    df = pd.read_csv(file_path)

    df_cleaned = df.dropna(subset=['status'])

    df_cleaned['value'] = df_cleaned['value'].fillna(0)
    return df_cleaned


if __name__ == "__main__":
    target_file = "data.csv"
    result = clean_data(target_file)
    print("\n[정제된 데이터]")
    print(result)

    result.to_csv("cleaned_data.csv", index=False)  
    print("\n[정제된 데이터가 'cleaned_data.csv'로 저장되었습니다.]")  
    