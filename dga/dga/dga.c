#include <stdio.h>
#include <string.h>
#include <time.h>
char* get_ip() {
	int flag;
	char ch;
	char *ip = (char *)malloc(20 * sizeof(char));
	FILE* fp, * f_dns;
	char tld[9][10] = { ".com", ".cn", ".net", ".org", ".com.cn", ".info", ".cc", ".top", ".wang" };
	int tld_num = 9;
	fp = fopen("LSTM.txt", "r+");
	f_dns = fopen("DNS.txt", "r+");
	if (fp == NULL || f_dns == NULL) {
		printf("文件打开时发生错误");
		return -1;
	}

	char domain[200][100];
	char dns_domain[100][100];
	int num = 0;
	int dns_num = 0;
	while (!feof(f_dns)) {
		fgets(dns_domain[dns_num], 100, (FILE*)f_dns);
		if (dns_domain[dns_num][strlen(dns_domain[dns_num]) - 1] == '\n'){
			printf("haa huanhang\n");
			dns_domain[dns_num][strlen(dns_domain[dns_num]) - 1] = 0;
		}
		dns_num++;
	}
	for (int i = 0; i < dns_num; i++) {
		printf("1:%s\n", dns_domain[i]);
	}
	while (!feof(fp)) {
		fgets(domain[num], 100, (FILE*)fp);
		if (domain[num][strlen(domain[num]) - 1] == '\n') {
			printf("haa huanhang\n");
			domain[num][strlen(domain[num]) - 1] = '\0';
		}
		//domain[num][strlen(domain[num])-1] = 0;
		
		num++;
	}
	for (int i = 0; i < num; i++) {
		printf("2:%s\n", domain[i]);
	}
	printf("%d\n", num);
	for (int i = 0; i < num; i++) {
		for (int j = 0; j < tld_num; j++) {
			char s[200];
			strcpy(s, domain[i]);
			printf("cpy:%s\n", s);
			strcat(s, tld[j]);
			printf("cat:%s\n", s);
			for (int k = 0; k < dns_num; k++) {
				printf("%s and %s\n", s, dns_domain[k]);
				if (strcmp(s, dns_domain[k]) == 0) {
					printf("yes\n");
					ip = s;
					printf("%s\n", ip);
					//return ip;
				}
				else
				{
					printf("no\n");
				}
			}
		}
		//rintf("%s\n", domain[i]);
	}
	fclose(fp);  //关闭文件，释放指向文件的指针
	return "no";
}

/*int main() {
	char* b = get_ip();
	printf("%s\n", b);
	return 0;
}*/

int main(void)
{
	long    i = 10000000L;
	clock_t start, finish;
	double  duration;
	/* 测量一个事件持续的时间*/
	printf("Time to do %ld empty loops is ", i);
	start = clock();
	while (i--);
	finish = clock();
	duration = (double)(finish - start) / CLOCKS_PER_SEC;
	printf("%f seconds\n", duration);
	system("pause");
}