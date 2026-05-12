import os
import pdfplumber
import re

def process_tyler_mowery_lectures(input_pdf, output_dir):
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    print(f"[*] Processing Tyler Mowery lectures from: {input_pdf}")
    
    try:
        with pdfplumber.open(input_pdf) as pdf:
            for i, page in enumerate(pdf.pages):
                content = page.extract_text()
                if not content: continue
                
                # 파일당 영상 하나 또는 페이지 하나로 매핑 (AX 준수)
                file_path = os.path.join(output_dir, f"Tyler_Mowery_Part_{i+1}.md")
                
                processed_content = f"**[Source: Tyler Mowery]**\n\n# Tyler Mowery Lecture Part {i+1}\n\n"
                
                # 타일러 고유 용어 강화 및 카테고리 태깅
                lines = content.split('\n')
                for line in lines:
                    line = line.strip()
                    if not line: continue
                    
                    # 태깅 로직 (Mowery의 핵심 키워드 중심)
                    if any(kw in line for kw in ["Belief", "Want", "Need", "Flaw", "신념", "결함"]):
                        line += " (Category: Character)"
                    elif any(kw in line for kw in ["Structure", "Beat", "Act", "구조", "플롯"]):
                        line += " (Category: Structure)"
                    elif any(kw in line for kw in ["Theme", "Moral", "Message", "테마", "도덕"]):
                        line += " (Category: Theme)"
                    
                    processed_content += line + "\n"
                
                # 핵심 원칙 요약 표 (AX 표준)
                processed_content += "\n## 핵심 주장 요약 (Tyler Mowery Principle)\n"
                processed_content += "| 원칙(Principle) | 세부 내용 |\n"
                processed_content += "| :--- | :--- |\n"
                processed_content += "| **Philosophy** | 캐릭터의 외부적 행동보다 내부적 신념(Belief)의 변화가 핵심 |\n"
                processed_content += "| **Logical Flow** | 모든 씬은 도덕적 논쟁(Moral Argument)을 강화해야 함 |\n"
                
                with open(file_path, "w", encoding="utf-8") as f:
                    f.write(processed_content)
                    
        print(f"[+] Extraction complete. Saved to {output_dir}")
    except Exception as e:
        print(f"[!] Error: {e}")

if __name__ == "__main__":
    # 절대 경로 및 glob 사용 (인코딩 문제 방지)
    base_path = r"c:\Users\psm23\Desktop\Autopus-NONE"
    raw_data_path = os.path.join(base_path, "projects", "NONE_시나리오", "raw_data")
    output_folder = os.path.join(base_path, "projects", "NONE_시나리오", "knowledge")
    
    import glob
    # '타일러' 또는 '.pdf' 패턴으로 파일 찾기
    pdf_files = glob.glob(os.path.join(raw_data_path, "*.pdf"))
    
    if pdf_files:
        # 타일러 모워리 파일 우선 선택
        target_file = next((f for f in pdf_files if "타일러" in f or "Tyler" in f), pdf_files[0])
        process_tyler_mowery_lectures(target_file, output_folder)
    else:
        print("[!] No PDF files found in raw_data.")
